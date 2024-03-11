import './App.css';
import React, { useState } from 'react';
import PackOrder from './components/PackOrder';

function App() {

  const [selectedSize, setSelectedSizes] = useState([5000, 2000, 1000, 500, 250]);
  const [sizes, setSizes] = useState([]);
  const [itemSize, setItemSize] = useState(0);

  const addOnChange = () => {
    setSelectedSizes(
      [parseInt(itemSize), ...selectedSize],
    );
    setSizes([parseInt(itemSize), ...sizes])
    setItemSize(0)
  }

  return (
    <div className="App">
      <div className="container">
        <h1><i class="bi bi-box"></i> Packaging Calculator <i class="bi bi-calculator-fill"></i></h1>
        <div className="row">
          <h2>Packages</h2>
          <form>
            <label>Size</label>
            <input
              id='re-input'
              aria-describedby="re-inputFeedback"
              type="number"
              value={itemSize}
              onChange={e => setItemSize(e.target.value)} s
            />
            <button
              type="button"
              className='btn btn-primary'
              onClick={addOnChange}><i class="bi bi-plus-square-fill"></i>
            </button>
          </form>
          <table className="table table-striped">
            <thead>
              <tr>
                <th>Size</th>
                <th>Options</th>
              </tr>
            </thead>
            <tbody>
              {selectedSize.map((size) => (
                <tr key={size}>
                  <td>
                    <i class="bi bi-box"></i> ({size})
                  </td>
                  <td>
                    <i class="bi bi-trash" onClick={() => setSelectedSizes(selectedSize.filter(s => s !== size))}></i>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <PackOrder sizes={selectedSize} />
      </div>
    </div>
  );
}

export default App;
