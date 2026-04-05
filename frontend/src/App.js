import React, { useState, useEffect } from 'react';
import axios from 'axios';

// Ensuring this matches your Go server port exactly
const API_URL = "http://localhost:3080";

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type':'application/json'
  }
})

const App = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [formData, setFormData] = useState({
    Title: '',
    Description: '',
    Price: '',
    ImgUrl: ''
  });

  // --- GET: Fetch all products from your Backend ---
  const fetchProducts = async () => {
    try {
      const response = await api.get(`${API_URL}/product`);
      setProducts(response.data);
      setLoading(false);
    } catch (error) {
      console.error("Fetch Error:", error);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  // --- POST: Send new product to your Backend ---
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const payload = { ...formData, Price: parseFloat(formData.Price) };
      // Fixed: Using API_URL instead of API_BASE
      await api.post(`${API_URL}/product-create`, payload);
      
      setFormData({ Title: '', Description: '', Price: '', ImgUrl: '' });
      fetchProducts(); // Refresh list
    } catch (error) {
      alert("Backend error: Check CORS or struct validation.");
    }
  };

  return (
    <div style={styles.container}>
      <header style={styles.header}>
        <h1 style={{ color: '#2c3e50' }}>Banana Store Console</h1>
        <p style={{ color: '#7f8c8d' }}>Connected to Go Backend on :3080</p>
      </header>

      {/* CREATE FORM */}
      <section style={styles.formSection}>
        <form onSubmit={handleSubmit} style={styles.form}>
          <input name="Title" placeholder="Title" value={formData.Title} onChange={handleInputChange} style={styles.input} required />
          <input name="Description" placeholder="Description" value={formData.Description} onChange={handleInputChange} style={styles.input} />
          <input name="Price" type="number" placeholder="Price" value={formData.Price} onChange={handleInputChange} style={styles.input} required />
          <input name="ImgUrl" placeholder="Image URL" value={formData.ImgUrl} onChange={handleInputChange} style={styles.input} required />
          <button type="submit" style={styles.button}>Add to Inventory</button>
        </form>
      </section>

      {/* PRODUCT LIST */}
      <div style={styles.grid}>
        {loading ? (
          <p>Loading...</p>
        ) : (
          products.map((item) => (
            <div key={item.ID} style={styles.card}>
              <img src={item.ImgUrl} alt={item.Title} style={styles.image} />
              <div style={{ padding: '15px' }}>
                <h3 style={{ margin: '0 0 10px 0' }}>{item.Title}</h3>
                <p style={{ color: '#27ae60', fontWeight: 'bold' }}>${item.Price}</p>
                <p style={{ fontSize: '12px', color: '#95a5a6' }}>{item.Description}</p>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

const styles = {
  container: { maxWidth: '1000px', margin: '0 auto', padding: '20px', fontFamily: 'sans-serif' },
  header: { textAlign: 'center', marginBottom: '30px' },
  formSection: { background: '#fff', padding: '20px', borderRadius: '8px', boxShadow: '0 2px 10px rgba(0,0,0,0.1)', marginBottom: '30px' },
  form: { display: 'flex', flexWrap: 'wrap', gap: '10px' },
  input: { flex: '1 1 200px', padding: '10px', borderRadius: '4px', border: '1px solid #ddd' },
  button: { width: '100%', padding: '12px', backgroundColor: '#f1c40f', border: 'none', borderRadius: '4px', fontWeight: 'bold', cursor: 'pointer' },
  grid: { display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))', gap: '20px' },
  card: { background: '#fff', borderRadius: '8px', overflow: 'hidden', boxShadow: '0 2px 5px rgba(0,0,0,0.1)' },
  image: { width: '100%', height: '150px', objectFit: 'cover' }
};

export default App;