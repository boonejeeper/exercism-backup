use time::{PrimitiveDateTime as DateTime, Duration};

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    // todo!("What time is a gigasecond later than {start}");
    // 1 000 000 000

    return start + (Duration::SECOND * 1000000000);
}
