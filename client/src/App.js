import React, { useState } from 'react';
import SearchPage from './pages/SearchPage';
import AddFoodPage from './pages/AddFoodPage';

function App() {
  const [page, setPage] = useState('search'); 

  return (
    <div className="App">
      <button onClick={() => setPage('search')}>Search Foods</button>
      <button onClick={() => setPage('add')}>Add Food</button>
      {page === 'search' ? <SearchPage /> : <AddFoodPage />}
    </div>
  );
}

export default App;
