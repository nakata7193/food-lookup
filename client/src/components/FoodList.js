import React, { useState, useEffect } from 'react';
import { searchFoods } from '../api/foodApi';

const FoodList = ({ searchTerm, addFoodToTable }) => {
  const [foods, setFoods] = useState([]);

  useEffect(() => {
    const fetchFoods = async () => {
      if (searchTerm) {
        const fetchedFoods = await searchFoods(searchTerm);
        setFoods(fetchedFoods);
      } else {
        setFoods([]);
      }
    };

    const timeoutId = setTimeout(() => {
      fetchFoods();
    }, 500);

    return () => clearTimeout(timeoutId);
  }, [searchTerm]);

  return (
    <div>
      <ul>
        {foods.map(food => (
          <li key={food.id} onClick={() => addFoodToTable(food)}>
            {food.name} - {food.calories} calories
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FoodList;

