import fetch from "node-fetch";

const data = {
  key: "104449841002335188382", // Put the redirected user's key here.
};

fetch("https://cs361micro-4qdz6le7kq-uc.a.run.app/userdata", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify(data),
})
  .then((response) => response.json())
  .then((data) => {
    console.log("data", data);
  })
  .catch((error) => {
    console.error("Error:", error);
  });
