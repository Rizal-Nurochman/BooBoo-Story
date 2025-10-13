import CreatorForm from '@/components/auth/CreatorForm'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/_authLayout/creator')({
  component: RouteComponent,
})

const borders=[
    {
        id:1,
        img:'/images/core/border.png',
        style:'w-[20%] absolute top-0 right-0'
    },
    {
        id:1,
        img:'/images/core/border.png',
        style:'w-[20%] rotate-[180deg] absolute bottom-0 left-0'
    },
]

function RouteComponent() {
  return(
    <div className='w-full h-screen relative'>
      {borders.map((img)=>(
          <img src={img.img} className={img.style} key={img.id}  />
      ))}
      <div className='w-full max-w-7xl mx-auto flex items-center pt-28 relative z-10'>
        <div className=''>
          {/* image */}
        </div>
        <div className="flex-1">
          <h1 className="text-5xl text-center font-bold text-primary">Tulis Kisah, Tebar Keajaiban</h1>
          <p className="text-center text-muted-foreground mt-3 text-xl max-w-xl mx-auto">
            Jadilah penulis yang membawa senyum dan keajaiban di setiap halaman. Dunia anak-anak menunggumu!
          </p>

          <CreatorForm />
        </div>
      </div>
    </div>
  )
}
