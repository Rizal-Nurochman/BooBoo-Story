import Image from 'next/image'
import React from 'react'
import { Button } from '../ui/button'
import Link from 'next/link'

const Header = () => {
  return (
    <header className='w-full absolute top-0 left-0 py-1 flex items-center justify-between px-8'>
        <Image src={'/images/logo.png'} width={100} height={100} alt='image logo' className='h-32 w-auto' />
        <Link href={'/login'} >
            <Button size={'lg'} className='cursor-pointer px-8 text-xl font-semibold py-3'>
                Login
            </Button>
        </Link>
    </header>
  )
}

export default Header