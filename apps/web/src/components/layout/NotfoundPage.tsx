import { useNavigate } from '@tanstack/react-router'

const NotfoundPage = () => {
    const navigate=useNavigate()
  return (
    <div className="flex pt-10 items-center justify-center min-h-screen bg-gradient-to-b from-sky-200 via-green-100 to-green-200">
      <div className="text-center px-6">
        <div className="mb-8 relative">
          <div className="flex justify-center gap-4 mb-4">
            <div className="w-16 h-24 bg-green-600 rounded-t-full relative">
              <div className="w-4 h-8 bg-amber-800 absolute bottom-0 left-1/2 -translate-x-1/2"></div>
            </div>
            <div className="w-20 h-28 bg-green-500 rounded-t-full relative -mt-4">
              <div className="w-5 h-10 bg-amber-700 absolute bottom-0 left-1/2 -translate-x-1/2"></div>
            </div>
            <div className="w-16 h-24 bg-green-600 rounded-t-full relative">
              <div className="w-4 h-8 bg-amber-800 absolute bottom-0 left-1/2 -translate-x-1/2"></div>
            </div>
          </div>
          
          <div className="text-8xl font-bold text-green-700 mb-4 drop-shadow-lg">
            404
          </div>
          
          <div className="flex justify-center mb-4">
            <div className="relative">
              <div className="absolute -top-8 left-2 w-3 h-10 bg-pink-300 rounded-full rotate-12"></div>
              <div className="absolute -top-8 right-2 w-3 h-10 bg-pink-300 rounded-full -rotate-12"></div>
              <div className="w-16 h-16 bg-pink-200 rounded-full relative">
                <div className="absolute top-5 left-4 w-2 h-2 bg-gray-800 rounded-full"></div>
                <div className="absolute top-5 right-4 w-2 h-2 bg-gray-800 rounded-full"></div>
                <div className="absolute top-8 left-1/2 -translate-x-1/2 w-1.5 h-1.5 bg-pink-400 rounded-full"></div>
              </div>
            </div>
          </div>
        </div>
        
        <h1 className="text-3xl font-bold text-green-700 mb-3">
          Oops! Kamu Tersesat di Hutan
        </h1>
        <p className="text-lg text-green-600 mb-6 max-w-md mx-auto">
          Halaman yang kamu cari tidak ditemukan. Kelinci kecil ini juga bingung mencarinya! 🌳
        </p>
        
        <button 
          onClick={() => navigate({ to: '..' })}
          className="bg-green-600 cursor-pointer hover:bg-green-700 text-white font-bold py-3 px-8 rounded-full shadow-lg transform hover:scale-105 transition-all duration-200"
        >
          🏠 Kembali ke Rumah
        </button>
        
        <div className="flex justify-center gap-2 mt-8">
          <div className="w-8 h-4 bg-green-400 rounded-t-full"></div>
          <div className="w-6 h-3 bg-green-500 rounded-t-full"></div>
          <div className="w-10 h-5 bg-green-400 rounded-t-full"></div>
          <div className="w-7 h-4 bg-green-500 rounded-t-full"></div>
          <div className="w-8 h-4 bg-green-400 rounded-t-full"></div>
        </div>
      </div>
    </div>
  )
}

export default NotfoundPage