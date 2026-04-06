import React from "react";

const ProductForm = ({ formData, onInputChange, onSubmit }) => {
  return (
    <form onSubmit={onSubmit} className="product-form">
      <input
        name="Title"
        placeholder="Title"
        value={formData.Title}
        onChange={onInputChange}
        className="product-form-input"
        required
      />
      <input
        name="Description"
        placeholder="Description"
        value={formData.Description}
        onChange={onInputChange}
        className="product-form-input"
      />
      <input
        name="Price"
        type="number"
        placeholder="Price"
        value={formData.Price}
        onChange={onInputChange}
        className="product-form-input"
        required
      />
      <input
        name="ImgUrl"
        placeholder="Image URL"
        value={formData.ImgUrl}
        onChange={onInputChange}
        className="product-form-input"
        required
      />
      <button type="submit" className="product-form-button">
        Add to Inventory
      </button>
    </form>
  );
};

export default ProductForm;
