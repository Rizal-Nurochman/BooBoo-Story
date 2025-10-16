import React from 'react'
import SearchHeader from './SearchHeader'
import UserDropdown from './UserDropdown'
import { Link } from '@tanstack/react-router'
import { Button } from '../ui/button'
import { Plus } from 'lucide-react'

const HeaderApp = () => {
  const [isScrolled, setIsScrolled] = React.useState(false)

  React.useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 0)
    }

    window.addEventListener('scroll', handleScroll)

    handleScroll()

    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  return (
    <header
      className={`sticky flex items-center justify-between top-0 left-0 w-full py-2 px-4 sm:px-6 lg:px-10 transition-all duration-300 z-20 ${
        isScrolled ? 'bg-primary/20 backdrop-blur-xl shadow-md' : 'bg-transparent'
      }`}
    >
      <SearchHeader />
      <div className='flex gap-4 items-center'>
        <Link to="/app/books/create">
          <Button size={'sm'} className="bg-gradient-to-r cursor-pointer from-green-500 to-emerald-600 text-white font-semibold shadow-md hover:shadow-lg hover:from-green-600 hover:to-emerald-700 transition-all duration-200">
            <Plus className="w-2 h-2 mr-2" />
            Buat Cerita
          </Button>
        </Link>
        <UserDropdown />
      </div>
    </header>
  )
}

export default HeaderApp
