import Category from "./ListCategory"
import Find from "./Find"

const SearchHeader = () => {
    const handleSearch = (query: string) => {
        console.log("Searching for:", query)
    }
  return (
    <div className="flex items-center gap-4 w-full max-w-md">
        <Category />
        <Find />
    </div>
  )
}

export default SearchHeader