import { APIAddress } from "./API_Address";
var nextLocation: number = 0;

export class Period {
  ID: number;
  Date_Stamp: number;
  Location: number;
  Duration: number;
  Type: string;
  Title: string;
  Description: string;
  constructor(
    ID: number,
    Date_Stamp: number,
    Location: number,
    Duration: number,
    Type: string,
    Title: string,
    Description: string
  ) {
    this.ID = ID;
    this.Date_Stamp = Date_Stamp;
    this.Location = Location;
    this.Duration = Duration;
    this.Type = Type;
    this.Title = Title;
    this.Description = Description;
  }
}

export async function postPeriods(
  periods: Array<Period>,
  p: Period,
  dateStamp: number
) {
  p.Date_Stamp = dateStamp;
  p.Location = nextLocation;
  p.ID = 0;
  const response = await fetch(APIAddress + "periods", {
    method: "POST",
    body: JSON.stringify(p),
    headers: { "Content-Type": "application/json" },
  });

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    let obj: Period = JSON.parse(await response.text());
    if (periods == null) periods = [obj];
    else periods.push(obj);
    nextLocation += 1;
  }
}

// note how this returns the array unlike the other functions
// this is because i am a discrace
export async function getPeriods(dateStamp: number): Promise<Array<Period>> {
  const response = await fetch(APIAddress + "periods/" + dateStamp);

  if (!response.ok) {
    console.error("Error");
    return [];
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
    return [];
  } else {
    let obj: Array<Period> = JSON.parse(await response.text());
    if (obj != null) {
      nextLocation = obj[obj.length - 1].Location + 1;
      return obj;
    } else {
      nextLocation = 0;
      return [];
    }
  }
}

export async function deletePeriods(periods: Array<Period>, pIndex: number) {
  const response = await fetch(APIAddress + "periods/" + periods[pIndex].ID, {
    method: "DELETE",
  });

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    periods.splice(pIndex, 1);
  }
}

export async function updatePeriods(
  periods: Array<Period>,
  p: Period,
  pIndex: number
) {
  const response = await fetch(APIAddress + "periods", {
    method: "PUT",
    body: JSON.stringify(p),
    headers: { "Content-Type": "application/json" },
  });

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    let obj: Period = JSON.parse(await response.text());
    if (periods != null) periods[pIndex] = obj;
  }
}
