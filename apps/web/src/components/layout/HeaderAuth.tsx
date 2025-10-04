import { Link } from "@tanstack/react-router"

const HeaderAuth = () => {
  return (
    <div className='relative w-full'>
       <div className="w-full max-w-[90%] absolute top-0 left-1/2 -translate-x-1/2 md:max-w-[85%] lg:max-w-[80%] px-4 md:px-6 lg:px-8 flex justify-between items-center">
          <Link to="/">
            <img src="/images/core/logo.png" alt="logo" className="w-81 px-8 h-auto" loading="lazy" />
          </Link>
       </div>
    </div>
  )
}

export default HeaderAuth