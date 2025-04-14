import axios from 'axios';


export const createOrder = async (orderData) => {
  try {
    const response = await axios.post('http://127.0.0.1:8082/api/order', orderData);
    return response.data;
  } catch (error) {
    console.error('Error creating order:', error);
    throw error;
  }
};
