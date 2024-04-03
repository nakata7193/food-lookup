export const searchFoods = async (query) => {
    const response = await fetch(`/api/foods?search=${query}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return await response.json();
  };
  


export const addFood = async (food) => {
    const response = await fetch('/api/food', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(food),
    });
  
    if (!response.ok) {
      throw new Error('Failed to add food item');
    }
    return response.json();
  };
  