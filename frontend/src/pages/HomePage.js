import React, { useEffect, useState } from "react";
import ProductForm from "../components/ProductForm";
import ProductGrid from "../components/ProductGrid";
import { createProduct, getProducts } from "../services/productService";

const HomePage = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState("");
  const [formData, setFormData] = useState({
    Title: "",
    Description: "",
    Price: "",
    ImgUrl: "",
  });

  const fetchProducts = async () => {
    try {
      const responseData = await getProducts();
      setProducts(responseData);
      setErrorMessage("");
      setLoading(false);
    } catch (error) {
      console.error("Fetch Error:", error);
      setErrorMessage("Could not load products. Make sure backend is running on port 3080.");
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

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const payload = { ...formData, Price: parseFloat(formData.Price) };
      await createProduct(payload);

      setFormData({ Title: "", Description: "", Price: "", ImgUrl: "" });
      fetchProducts();
    } catch (error) {
      alert("Backend error: Check CORS or struct validation.");
    }
  };

  return (
    <div className="app-page-wrap">
      <div className="app-container">
        <header className="app-base-card app-header-card">
          <h1 className="app-title">Banana Store Console</h1>
          <p className="app-subtitle">Connected to Go Backend on :3080</p>
        </header>

        <section className="app-base-card">
          <h2 className="app-section-title">Add Product</h2>
          <ProductForm
            formData={formData}
            onInputChange={handleInputChange}
            onSubmit={handleSubmit}
          />
        </section>

        {errorMessage ? <p className="app-error-text">{errorMessage}</p> : null}

        <section className="app-base-card">
          <h2 className="app-section-title">Inventory</h2>
          <ProductGrid loading={loading} products={products} />
        </section>
      </div>
    </div>
  );
};

export default HomePage;
