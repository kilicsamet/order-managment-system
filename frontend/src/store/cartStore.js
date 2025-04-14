import { create } from 'zustand';
import { getInventory, getProduct } from '../api/inventory';
import { createOrder } from '../api/order';

const useCartStore = create((set) => ({
  items: [],
  totalPrice: 0,
  isLoading: false,
  isEmpty: false,
  successMessage: '',
  errorMessage: '',

  createOrder: async () => {
    set({ isLoading: true });

    try {
      const orderItems = useCartStore.getState().items.map(item => ({
        product_id: item.id,
        quantity: item.quantity
      }));

      const orderData = {
        status: 'pending',
        order_items: orderItems
      };
      const result = await createOrder(orderData);

      if (result) {
        console.log('Sipariş başarıyla oluşturuldu:', result);
        set({ items: [], successMessage: 'Sipariş başarıyla oluşturuldu. Aynı ürünlerle tekrar satış yapmak istiyorusanız sayfayı yenileyin!', errorMessage:'' });
      }
      set({ isLoading: false });
    } catch (error) {
      console.error('Bir hata oluştu:', error);
      set({
        isLoading: false,
        errorMessage: error.response?.data?.error || 'Bir hata oluştu, lütfen tekrar deneyin.',
        successMessage: ''
      });
    }
  },

  fetchProducts: async () => {
    set({ isLoading: true });
    try {
      const [responseInventory, responseProduct] = await Promise.all([
        getInventory(),
        getProduct()
      ]);

      const productsData = responseInventory.inventory_products.map(inventoryProduct => {
        const productDetails = responseProduct.products.find(product => product.id === inventoryProduct.product_id);

        return {
          ...inventoryProduct,
          ...productDetails,
          quantity: 1
        };
      });

      set({
        items: productsData,
        isLoading: false,
        isEmpty: productsData.length === 0
      });
    } catch (error) {
      console.error('Ürün verisi alınamadı:', error);
      set({ isLoading: false });
    }
  },
  increaseQuantity: (productId) => set((state) => {
    const updatedItems = state.items.map(item =>
      item.id === productId ? { ...item, quantity: item.quantity + 1 } : item
    );
    return { items: updatedItems };
  }),

  decreaseQuantity: (productId) => set((state) => {
    const updatedItems = state.items.map(item =>
      item.id === productId && item.quantity > 1
        ? { ...item, quantity: item.quantity - 1 }
        : item
    );
    return { items: updatedItems };
  }),

  calculateTotal: () => set((state) => {
    const totalPrice = state.items.reduce((total, item) => total + item.price * item.quantity, 0);
    return { totalPrice };
  }),

  checkIfEmpty: () => set((state) => {
    return { isEmpty: state.items.length === 0 };
  }),

  setLoading: (isLoading) => set({ isLoading }),

  removeItem: (id) => set((state) => ({
    items: state.items.filter((item) => item.id !== id),
  })),
}));


export default useCartStore;
