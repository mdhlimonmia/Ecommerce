import { API_URL } from "../config/api";

const DEFAULT_HEADERS = {
  "Content-Type": "application/json",
};

const parseJsonResponse = async (response) => {
  if (!response.ok) {
    const text = await response.text();
    throw new Error(text || `Request failed with status ${response.status}`);
  }

  return response.json();
};

export const getProducts = async () => {
  const response = await fetch(`${API_URL}/products`, {
    method: "GET",
    headers: DEFAULT_HEADERS,
  });

  return parseJsonResponse(response);
};

export const createProduct = async (product) => {
  const response = await fetch(`${API_URL}/products?user=admin&password=1234`, {
    method: "POST",
    headers: DEFAULT_HEADERS,
    body: JSON.stringify(product),
  });

  return parseJsonResponse(response);
};

export const getProductById = async (productId) => {
  const response = await fetch(`${API_URL}/products/${productId}`, {
    method: "GET",
    headers: DEFAULT_HEADERS,
  });

  return parseJsonResponse(response);
};
