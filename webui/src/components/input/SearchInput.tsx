import { search as searchAPI } from '@/apis/search.api'
import { END_POINT_MINIO } from '@/constants/const'
import useDebounce from '@/hooks/useDebounce'
import { useQuery } from '@tanstack/react-query'
import { Avatar } from 'primereact/avatar'
import { Card } from 'primereact/card'
import { useState } from 'react'
// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface SearchInputProps {}
// eslint-disable-next-line no-empty-pattern
const SearchInput: React.FC<SearchInputProps> = ({}) => {
  const [search, setSearch] = useState<string>('')
  const debouncedSearch = useDebounce(search, 500)
  const searchesRes = useQuery({
    queryKey: ['search', debouncedSearch],
    queryFn: () => {
      if (debouncedSearch) return searchAPI(debouncedSearch)
      return null
    }
  })
  const searches = searchesRes?.data?.data?.searches
  return (
    <div className='relative'>
      <div className='flex items-center justify-between space-x-1 '>
        <div className='relative bg-gray-100 dark:bg-gray-600 p-5 w-full h-11 dark:text-gray-300 rounded-full flex items-center justify-center cursor-pointer'>
          <i className='pi pi-search mr-2'></i>
          <input
            type='text'
            placeholder='Search'
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className='outline-none bg-transparent hidden xl:inline-block'
          />
        </div>
      </div>
      {searches && searches?.length > 0 && (
        <Card className='absolute top-12'>
          <ul>
            {searches.map((search) => {
              return (
                <li key={`search_feed_${search?.id}`} className='hover:bg-gray-200 p-3 rounded-lg'>
                  <a
                    href={search?.name ? `/groups/${search?.id}` : `/profile?id=${search.id}`}
                    className='flex w-96 items-center gap-x-3'
                  >
                    <i className='pi pi-search'></i>
                    <div className='flex items-center justify-start w-full gap-x-2'>
                      <Avatar
                        image={END_POINT_MINIO + (search?.avatarUrl ? search?.avatarUrl : search?.profileUrl)}
                        size='normal'
                        shape='circle'
                      />
                      <span className='font-medium text-base'>{search.name ? search.name : search.fullName}</span>
                      <span className='text-xs italic'>{search.name ? '· Group' : '· ' + search.position}</span>
                    </div>
                  </a>
                </li>
              )
            })}
          </ul>
        </Card>
      )}
    </div>
  )
}
export default SearchInput
