import React, { useState } from "react";

export const CbApp = () => {
  // State with list of all checked item
  const [checked, setChecked] = useState([]);
  const checkList = [
    {
        "Name":"Apple",
        "score":"35"
    },
    { 
        "Name":"Banana",
        "score":"5"
    },
    { 
        "Name":"Tea",
        "score":"25"
    },
    { 
        "Name":"Coffee",
        "score":"15"
    }];

  // Add/Remove checked item from list
  const handleCheck = (event) => {
    var updatedList = [...checked];
    if (event.target.checked) {
      updatedList = [...checked, event.target.value];
    } else {
      updatedList.splice(checked.indexOf(event.target.value), 1);
    }
    setChecked(updatedList);
  };

  // Generate string of checked items
  const checkedItems = checked.length
    ? checked.reduce((total, item) => {
        return total + ", " + item;
      })
    : "";

  // Return classes based on whether item is checked
  var isChecked = (item) =>
    checked.includes(item) ? "checked-item" : "not-checked-item";

  return (
    <div className="app">
      <div className="checkList">
        <div className="title">Your CheckList:</div>
        <div className="list-container">
          {checkList.map((item, index) => (
            <div key={index}>
              <input value={item.score} type="checkbox" onChange={handleCheck} />
              <span className={isChecked(item.Name)}>{item.Name}</span>
            </div>
          ))}
        </div>
      </div>

      <div>
        {`Items checked are: ${checkedItems}`}
      </div>
    </div>
  );
}