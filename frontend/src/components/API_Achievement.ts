import { APIAddress } from "./API_Address";

export class Achievement {
  ID: number;
  Date_Stamp: number;
  Text: string;
  constructor(ID: number, Date_Stamp: number, Text: string) {
    this.ID = ID;
    this.Date_Stamp = Date_Stamp;
    this.Text = Text;
  }
}

export async function postAchievements(
  achievements: Array<Achievement>,
  a: Achievement,
  dateStamp: number
) {
  a.Date_Stamp = dateStamp;
  a.ID = 0;
  const response = await fetch(APIAddress + "achievements", {
    method: "POST",
    body: JSON.stringify(a),
    headers: { "Content-Type": "application/json" },
  });

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    let obj: Achievement = JSON.parse(await response.text());
    if (achievements == null) achievements = [obj];
    else achievements.push(obj);
  }
}

// note how this returns the array unlike the other functions
// this is because i am a discrace
export async function getAchievements(
  dateStamp: number
): Promise<Array<Achievement>> {
  const response = await fetch(APIAddress + "achievements/" + dateStamp);

  if (!response.ok) {
    console.error("Error");
    return [];
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
    return [];
  } else {
    let obj: Array<Achievement> = JSON.parse(await response.text());
    if (obj != null) {
      return obj;
    } else {
      return [];
    }
  }
}

export async function deleteAchievements(
  achievements: Array<Achievement>,
  index: number
) {
  const response = await fetch(
    APIAddress + "achievements/" + achievements[index].ID,
    {
      method: "DELETE",
    }
  );

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    achievements.splice(index, 1);
  }
}

export async function updateAchievements(
  achievements: Array<Achievement>,
  p: Achievement,
  pIndex: number
) {
  const response = await fetch(APIAddress + "achievements", {
    method: "PUT",
    body: JSON.stringify(p),
    headers: { "Content-Type": "application/json" },
  });

  if (!response.ok) {
    console.error("Error");
  } else if (response.status >= 400) {
    console.error("HTTP Error: " + response.status + " - " + response.status);
  } else {
    let obj: Achievement = JSON.parse(await response.text());
    if (achievements != null) achievements[pIndex] = obj;
  }
}
