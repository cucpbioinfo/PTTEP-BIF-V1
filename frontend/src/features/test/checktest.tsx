import React, { Fragment, useState, useEffect } from "react";

export const CheckBoxTest = () => {
  const [person, setPerson] = useState([]);

  const CheckHandler = (e) => {
    console.log(e.target.value);
    if(person.includes(e.target.value))
    {
      setPerson((prev) => [...prev]);
      console.log("already have ",e.target.value);
      
      //const numbers = [1, 2, 3, 3, 3, 3, 2, 1]; // replace with person
      //const target = 3; // replace with person e.target.value
      var i = 0;
      while (i < person.length) {
        if (person[i] === e.target.value) {
          person.splice(i, 1);
        } else {
            ++i;
        }
      }
      //console.log(person); 



    }
    else
    {
      setPerson((prev) => [...prev, e.target.value]);
    }
  };

  useEffect(() => {
    //Print data each time the checkbox is "checked" or "unchecked"
    console.log(person);
  }, [person]);

  return (
    <Fragment>
      <input
        type="checkbox"
        id="jane"
        name="jane"
        value="jane"
        onClick={CheckHandler}
      />
      <label htmlFor="jane">jane</label>
      <br />

      <input
        type="checkbox"
        id="Mike"
        name="Mike"
        value="Mike"
        onClick={CheckHandler}
      />
      <label htmlFor="Mike">Mike</label>
      <br />

      <input
        type="checkbox"
        id="board"
        name="board"
        value="board"
        onClick={CheckHandler}
      />
      <label htmlFor="board">board</label>
      <br />
    </Fragment>
  )
}