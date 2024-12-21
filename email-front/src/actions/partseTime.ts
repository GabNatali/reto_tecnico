import type { TimeUnit } from "@/interfaces";

const unitToMicroseconds: Record<TimeUnit, number> = {
    minute: 60 * 1000 * 1000,
    hour: 60 * 60 * 1000 * 1000,
    day: 24 * 60 * 60 * 1000 * 1000,
    week: 7 * 24 * 60 * 60 * 1000 * 1000,
};
export const parseTime = (amount:number , unit: TimeUnit) => {
  const nowMicroseconds = Date.now() * 1000;

  const timeSend = amount * unitToMicroseconds[unit];

  const endTime = nowMicroseconds;
  const startTime = endTime - timeSend;

  return{
    endTime,
    startTime
  }

}
