import React from 'react';
import AddFoodForm from '../components/AddFoodForm';
import { addFood } from '../api/foodApi'; 

const AddFoodPage = () => {
  const handleSubmit = async (food) => {
    try {
      await addFood(food);
      alert('Food added successfully!');
    } catch (error) {
      alert('Failed to add food');
      console.error(error);
    }
  };

  return (
    <div>
      <h1>Add New Food Item</h1>
      <AddFoodForm onSubmit={handleSubmit} />
    </div>
  );
};

export default AddFoodPage;
