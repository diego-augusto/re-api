import './App.css';
import React, { useState } from 'react';
import PackOrder from './components/PackOrder';

function App() {

  const [selectedSize, setSelectedSizes] = useState([5000, 2000, 1000, 500, 250]);
  const [sizes, setSizes] = useState([]);
  const [itemSize, setItemSize] = useState(0);

  const addOnChange = () => {
    setSelectedSizes(
      [...selectedSize, parseInt(itemSize)],
    );
    setSizes([...sizes, parseInt(itemSize)])
    setItemSize(0)
  }

  return (
    <div className="App">
      <h1>Package Calculator</h1>
      <div className="container">
      <div className="row">
          <h2>Package Sizes</h2>
          <form>
            <label>Package Size</label>
            <input type="number" value={itemSize} onChange={e => setItemSize(e.target.value)} />
            <button type="button" className='btn btn-primary' onClick={addOnChange}>Add New Package Size</button>
          </form>
          <table>
            <thead>
              <tr>
                <th>Package Size</th>
              </tr>
            </thead>
            <tbody>
              {selectedSize.map((size) => (
                <tr key={size}>
                  <td>{size}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <PackOrder sizes={selectedSize}/>
      </div>
    </div>
  );
}

export default App;
