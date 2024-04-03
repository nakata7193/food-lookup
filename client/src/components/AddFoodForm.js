import React, { useState } from 'react';

const AddFoodForm = ({ onSubmit }) => {
  const [food, setFood] = useState({
    name: '',
    calories: 0,
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFood(prevFood => ({
      ...prevFood,
      [name]: name === 'calories' ? parseInt(value, 10) : value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(food);
    setFood({ name: '', calories: 0 });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="name">Food Name:</label>
        <input
          type="text"
          id="name"
          name="name"
          value={food.name}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="calories">Calories:</label>
        <input
          type="number"
          id="calories"
          name="calories"
          value={food.calories}
          onChange={handleChange}
          required
        />
      </div>
      <button type="submit">Add Food</button>
    </form>
  );
};

export default AddFoodForm;
