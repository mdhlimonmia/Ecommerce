import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { getProductById } from "../services/productService";

const ProductDetailsPage = () => {
  const { id } = useParams();
  const [product, setProduct] = useState(null);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const data = await getProductById(id);
        setProduct(data);
        setErrorMessage("");
      } catch (error) {
        setErrorMessage("Product not found or backend is unavailable.");
      } finally {
        setLoading(false);
      }
    };

    fetchProduct();
  }, [id]);

  return (
    <div className="app-page-wrap">
      <div className="app-container">
        <section className="app-base-card">
          <Link to="/" className="product-back-link">
            &larr; Back to products
          </Link>

          {loading ? (
            <p>Loading product...</p>
          ) : errorMessage ? (
            <p className="app-error-text">{errorMessage}</p>
          ) : (
            <div className="product-details-layout">
              <img
                src={product.ImgUrl}
                alt={product.Title}
                className="product-details-image"
              />
              <div>
                <h1 className="product-details-title">{product.Title}</h1>
                <p className="product-details-price">${product.Price}</p>
                <p className="product-details-description">{product.Description}</p>
                <p className="product-details-id">Product ID: {product.ID}</p>
              </div>
            </div>
          )}
        </section>
      </div>
    </div>
  );
};

export default ProductDetailsPage;
