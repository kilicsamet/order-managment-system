import axios from 'axios';

export const getInventory = async () => {
  try {
    const response = await axios.get('http://127.0.0.1:8081/api/inventory_products');
    return response.data;
  } catch (error) {
    console.error('Error fetching inventory:', error);
    throw error;
  }
};
export const getProduct = async () => {
    try {
      const response = await axios.get('http://127.0.0.1:8080/api/products');
      return response.data;
    } catch (error) {
      console.error('Error fetching inventory:', error);
      throw error;
    }
  };

export const addProductInventory = async (productData) => {
  try {
    const response = await axios.post('http://127.0.0.1:8081/api/inventory_product', productData);
    return response.data;
  } catch (error) {
    console.error('Error adding product:', error);
    throw error;
  }
};
export const addProduct = async (productData) => {
    try {
      const response = await axios.post('http://127.0.0.1:8080/api/product', productData);
      return response.data;
    } catch (error) {
      console.error('Error adding product:', error);
      throw error;
    }
  };
