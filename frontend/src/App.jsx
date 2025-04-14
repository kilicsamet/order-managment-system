import { Toaster } from "react-hot-toast";
import "./App.css";
import CartPage from "./pages/CartPage";

function App() {
  return (
    <div className="flex justify-between items-start">
      <Toaster
        position="top-right"
        toastOptions={{
          style: {
            fontSize: "25px",
            padding: "12px 16px",
            minWidth: "250px",
          },
          success: {
            style: {
              background: "#daf5d4",
              color: "#2b6b2f",
            },
          },
          error: {
            style: {
              background: "#ffe5e5",
              color: "#a30000",
            },
          },
        }}
      />

      <CartPage />
    </div>
  );
}

export default App;
