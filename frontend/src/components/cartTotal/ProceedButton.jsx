import React from 'react'

const ProceedButton = ({isLoading,setLoading, createOrder}) => {
  return (
    <div className="mt-8 w-full flex justify-center">
    <button
      onClick={() => {
        setLoading(true);
        setTimeout(() => {
        createOrder()
        }, 2000);
      }}
      className="w-full px-6 py-2 bg-green-950 text-white rounded-md disabled:bg-gray-300"
      disabled={isLoading}
    >
      Proceed to Checkout
    </button>
  </div>
  )
}

export default ProceedButton
