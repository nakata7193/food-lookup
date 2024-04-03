import React, { useState } from 'react';
import SearchBar from '../components/SearchBar';
import FoodList from '../components/FoodList';
import FoodTable from '../components/FoodTable';

const SearchPage = () => {
  const [selectedFoods, setSelectedFoods] = useState([]);

  const addFoodToTable = (food) => {
    const isPresent = selectedFoods.some(selectedFood => selectedFood.id === food.id);
    if (!isPresent) {
      setSelectedFoods(prevFoods => [...prevFoods, food]);
    }
  };

  return (
    <div>
      <h1>Search for Foods</h1>
      <SearchBar />
      <FoodList addFoodToTable={addFoodToTable} />
      <FoodTable foods={selectedFoods} />
    </div>
  );
};

export default SearchPage;
