import React, {useState } from 'react';

function PackOrder({sizes}) {

  const [loading, setLoading] = useState(false);
  const [item, setItem] = useState(0);
  const [packs, setPacks] = useState([]);

  const handleClick = async () => {
    try {
      setLoading(true)
      const data = await (await fetch(`${process.env.REACT_APP_API_URL}/pack`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({items: parseInt(item), sizes: sizes})
      })).json()
      setPacks(data)
    } catch (err) {
      console.log(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
        <div className='row'>
          <h2>Package Details</h2>
          <form>
            <label>Items</label>
            <input type="number" value={item} onChange={e => setItem(e.target.value)} />
            <button type="button" className='btn btn-primary' onClick={handleClick}>Get Packages</button>
          </form>
          <table>
            <thead>
              <tr>
                <th>Package</th>
                <th>Quantity</th>
              </tr>
            </thead>
            <tbody>
              {loading ? <tr><td>loading...</td></tr> : packs.map((pack) => (
                <tr key={pack.size}>
                  <td>{pack.size}</td>
                  <td>{pack.quantity}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
  );
}

export default PackOrder;
