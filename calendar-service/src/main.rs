mod encryptor;
mod service;

use std::env;
use tonic::{transport::Server, Request, Response, Status};

use calendar::calendar_service_server::{CalendarService as Calendar, CalendarServiceServer};
use calendar::{
    CreateCalendarRequest, CreateCalendarResponse, GetCalendarRequest, GetCalendarResponse,
};

pub mod calendar {
    tonic::include_proto!("calendar");
}

#[derive(Debug, Default)]
pub struct CalendarService {}

#[tonic::async_trait]
impl Calendar for CalendarService {
    // TODO: implement full method
    async fn get_calendar(
        &self,
        request: Request<GetCalendarRequest>,
    ) -> Result<Response<GetCalendarResponse>, Status> {
        let request_iv = request.get_ref().calendar_iv.as_bytes();
        if request_iv.len() != 12 {
            return Err(Status::invalid_argument("Invalid calendar_iv length"));
        }
        let mut iv: [u8; 12] = Default::default();
        iv.copy_from_slice(&request_iv[0..12]);

        let file = service::calendar::read(request.get_ref().calendar_id.as_str(), &iv);
        if file.is_err() {
            let reply = calendar::GetCalendarResponse {
                error: true,
                message: file.err().unwrap().to_string(),
                calendar: Vec::<u8>::new(),
            };
            Ok(Response::new(reply))
        } else {
            // TODO: Decrypt file

            let reply = calendar::GetCalendarResponse {
                error: false,
                message: "Calendar retrieved".to_string(),
                calendar: file.unwrap().as_bytes().to_vec(),
            };
            Ok(Response::new(reply))
        }
    }

    async fn create_calendar(
        &self,
        request: Request<CreateCalendarRequest>,
    ) -> Result<Response<CreateCalendarResponse>, Status> {
        let request_iv = request.get_ref().calendar_iv.as_bytes();
        if request_iv.len() != 12 {
            return Err(Status::invalid_argument("Invalid calendar_iv length"));
        }
        let mut iv: [u8; 12] = Default::default();
        iv.copy_from_slice(&request_iv[0..12]);

        let calendar_ics = service::calendar::create(request.get_ref().calendar_entities.clone());

        let write_check = service::calendar::write(
            calendar_ics.to_string(),
            request.get_ref().calendar_id.as_str(),
            &iv,
        );

        match write_check {
            Ok(_) => {
                let reply = calendar::CreateCalendarResponse {
                    error: false,
                    message: "Calendar created".to_string(),
                };
                Ok(Response::new(reply))
            }
            Err(msg) => {
                let reply = calendar::CreateCalendarResponse {
                    error: true,
                    message: msg.to_string(),
                };
                Ok(Response::new(reply))
            }
        }
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("Starting calendars server...");
    // TODO: Set default values for env vars
    let grpc_port = env::var("GRPC_PORT").expect("Expected GRPC_PORT to be set");
    let addr = format!("0.0.0.0:{}", grpc_port).parse()?;
    let calendar_service = CalendarService::default();

    Server::builder()
        .add_service(CalendarServiceServer::new(calendar_service))
        .serve(addr)
        .await?;

    Ok(())
}
