import React, { useState } from "react";
import axios from "axios";

function Button(props) {
  const handleInsert = () => {
    const data = { id: 16, name: "React", mobile: "8080" };
    axios
      .post("http://localhost:8080/addusers", data)
      .then((response) => {
        console.log(response.data); // handle response from server
      })
      .catch((error) => {
        console.error(error); // handle error
      });
  };

  return (
    <div>
      <button onClick={handleInsert}>Update</button>
    </div>
  );
}

export default Button;
