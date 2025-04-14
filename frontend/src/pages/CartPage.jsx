import React, { useEffect } from "react";
import useCartStore from "../store/cartStore";
import Breadcrumb from "../components/Breadcrumb";
import Loading from "../components/Loading";
import IsEmty from "../components/IsEmty";
import Cart from "../components/Cart";
import toast from "react-hot-toast";

const CartPage = () => {
  const {
    fetchProducts,
    items,
    totalPrice,
    increaseQuantity,
    decreaseQuantity,
    calculateTotal,
    isEmpty,
    setLoading,
    isLoading,
    removeItem,
    createOrder,
    successMessage,
    errorMessage
  } = useCartStore();
  useEffect(() => {
    fetchProducts();
  }, [fetchProducts]);
  useEffect(() => {
    calculateTotal();
  }, [items, calculateTotal]);
  useEffect(() => {
    if (successMessage) {
      toast.success(successMessage);
      useCartStore.setState({ successMessage: '' });
    }
  }, [successMessage]);

  useEffect(() => {
    if (errorMessage) {
      toast.error(errorMessage);
      useCartStore.setState({ errorMessage: '' });
    }
  }, [errorMessage]);

  if (isEmpty) {
    return <IsEmty />;
  }

  return (
    <div className="w-full lg:w-3/4 px-4 py-6 mx-auto">
      <Breadcrumb />
      {isLoading ? (
        <Loading />
      ) : items.length === 0 ? (
        <IsEmty />
      ) : (
        <Cart
        createOrder={createOrder}
        items={items}
        increaseQuantity={increaseQuantity}
        decreaseQuantity={decreaseQuantity}
        removeItem={removeItem}
        totalPrice={totalPrice}
        setLoading={setLoading}
        isLoading={isLoading}
      />
      )
      }
    </div>
  );
};

export default CartPage;
