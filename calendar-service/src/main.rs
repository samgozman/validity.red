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
    // ReadCalendar handler
    // 1. Read file with ID (from request data) as filename
    // 2. Decode file with AES-256
    // - optional: check if file is valid (if possible)
    // 3. Send decoded structure as a gRPC stream to the broker
    // TODO: implement full method
    async fn get_calendar(
        &self,
        request: Request<GetCalendarRequest>,
    ) -> Result<Response<GetCalendarResponse>, Status> {
        println!("Got a request: {:?}", request);
        let reply = calendar::GetCalendarResponse {
            error: false,
            message: "Not implemented".to_string(),
            calendar: Vec::<u8>::new(),
        };
        Ok(Response::new(reply))
    }

    // CreateCalendar handler
    // 1. Recieve combined calendar data from broker
    // 2. Create .ics structure from data
    // 3. Write .ics to a buffer
    // 4. Encode buffer with AES-256
    // 5. Write encoded buffer to file with ID (from request data) as filename
    // 7. Replace previous calendar file (if exists)
    // TODO: implement full method
    async fn create_calendar(
        &self,
        request: Request<CreateCalendarRequest>,
    ) -> Result<Response<CreateCalendarResponse>, Status> {
        println!("Got a request: {:?}", request);
        let reply = calendar::CreateCalendarResponse {
            error: false,
            message: "Not implemented".to_string(),
        };
        Ok(Response::new(reply))
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
