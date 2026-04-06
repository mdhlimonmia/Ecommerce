import React from "react";
import { Link } from "react-router-dom";

const ProductGrid = ({ loading, products }) => {
  return (
    <div className="product-grid">
      {loading ? (
        <p>Loading...</p>
      ) : (
        products.map((item) => (
          <div key={item.ID} className="product-grid-card">
            <img src={item.ImgUrl} alt={item.Title} className="product-grid-image" />
            <div className="product-grid-content">
              <h3 className="product-grid-name">{item.Title}</h3>
              <p className="product-grid-price">${item.Price}</p>
              <p className="product-grid-description">{item.Description}</p>
              <Link to={`/products/${item.ID}`} className="product-grid-link">
                View Details
              </Link>
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default ProductGrid;
