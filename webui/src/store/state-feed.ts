import { create } from 'zustand'
import { devtools } from 'zustand/middleware'

type StateFeedStore = {
    typeFeed: string
    setTypeFeed: (type: string) => void
}
export const useStateFeedStore = create<StateFeedStore>()(
    devtools(
        (set) => ({
            typeFeed: "default",
            setTypeFeed: (type) => set({ typeFeed: type}),
        }),
        {
            enabled: true,
            name: 'state-feed-store'
        }
    )
)
