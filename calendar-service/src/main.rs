mod encryptor;
mod service;

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
        let file = service::calendar::read(request.get_ref().calendar_id.as_str());
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
        let calendar_ics = service::calendar::create(request.get_ref().calendar_entities.clone());
        // TODO: encrypt calendar_ics

        let write_check = service::calendar::write(
            calendar_ics.to_string(),
            request.get_ref().calendar_id.as_str(),
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
    // TODO: add from ENV
    let addr = "[::1]:50051".parse()?;
    let calendar_service = CalendarService::default();

    Server::builder()
        .add_service(CalendarServiceServer::new(calendar_service))
        .serve(addr)
        .await?;

    Ok(())
}
