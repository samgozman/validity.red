mod encryptor;
mod service;

use chrono_tz::Tz;
use std::env;
use tonic::{transport::Server, Request, Response, Status};

use calendar::calendar_service_server::{CalendarService as Calendar, CalendarServiceServer};
use calendar::{CreateCalendarRequest, GetCalendarRequest, GetCalendarResponse};

pub mod calendar {
    tonic::include_proto!("calendar");
}

#[derive(Debug, Default)]
pub struct CalendarService {}

#[tonic::async_trait]
impl Calendar for CalendarService {
    async fn get_calendar(
        &self,
        request: Request<GetCalendarRequest>,
    ) -> Result<Response<GetCalendarResponse>, Status> {
        let request_iv = request.get_ref().calendar_iv.clone();
        if request_iv.len() != 12 {
            return Err(Status::invalid_argument("Invalid calendar_iv length"));
        }
        let mut iv: [u8; 12] = Default::default();
        iv.copy_from_slice(&request_iv[0..12]);

        let file = service::calendar::read(request.get_ref().calendar_id.as_str(), &iv);
        if file.is_err() {
            Err(file.err().unwrap())
        } else {
            let reply = calendar::GetCalendarResponse {
                calendar: file.unwrap().as_bytes().to_vec(),
            };
            Ok(Response::new(reply))
        }
    }

    async fn create_calendar(
        &self,
        request: Request<CreateCalendarRequest>,
    ) -> Result<Response<()>, Status> {
        let request_iv = request.get_ref().calendar_iv.clone();
        if request_iv.len() != 12 {
            return Err(Status::invalid_argument("Invalid calendar_iv length"));
        }
        let mut iv: [u8; 12] = Default::default();
        iv.copy_from_slice(&request_iv[0..12]);

        let timezone_input = &request.get_ref().timezone.parse::<Tz>();
        let timezone: Tz;
        if timezone_input.is_err() {
            return Err(Status::invalid_argument("Invalid timezone value"));
        } else {
            timezone = timezone_input.as_ref().unwrap().clone();
        }

        let calendar_ics =
            service::calendar::create(&request.get_ref().calendar_entities.clone(), timezone);

        let write_check = service::calendar::write(
            calendar_ics.to_string(),
            request.get_ref().calendar_id.as_str(),
            &iv,
        );
        match write_check {
            Ok(_) => Ok(Response::new(())),
            Err(msg) => Err(Status::internal(msg.to_string())),
        }
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Sentry integration
    let _guard = sentry::init((
        env::var("SENTRY_DSN").expect("Expected SENTRY_DSN to be set"),
        sentry::ClientOptions {
            release: sentry::release_name!(),
            sample_rate: 1.0,
            traces_sample_rate: 0.2,
            ..Default::default()
        },
    ));

    println!("Starting calendars server...");
    let grpc_port = env::var("GRPC_PORT").expect("Expected GRPC_PORT to be set");
    let addr = format!("0.0.0.0:{}", grpc_port).parse()?;
    let calendar_service = CalendarService::default();

    Server::builder()
        .add_service(CalendarServiceServer::new(calendar_service))
        .serve(addr)
        .await?;

    Ok(())
}
