import React from 'react';

const Loading = () => {
  return (
    <div className="flex items-center justify-center min-h-100">
      <div className="flex items-center space-x-2 p-6 bg-transparent">
        <div className="w-20 h-20 border-t-4 border-blue-500 border-solid rounded-full animate-spin"></div>
      </div>
    </div>
  );
}

export default Loading;
