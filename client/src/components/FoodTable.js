import React from 'react';

const FoodTable = ({ foods }) => {
  const totalCalories = foods.reduce((acc, food) => acc + food.calories, 0);

  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Calories</th>
          </tr>
        </thead>
        <tbody>
          {foods.map(food => (
            <tr key={food.id}>
              <td>{food.name}</td>
              <td>{food.calories}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <div>Total Calories: {totalCalories}</div>
    </div>
  );
};

export default FoodTable;
