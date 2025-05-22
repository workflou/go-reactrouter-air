import { Outlet } from "react-router";
import useSWR from "swr"
import { fetcher } from "../lib/api";
import SetupForm from "../components/setup/SetupForm";

export default function AppLayout() {
    const { data, isLoading, error } = useSWR('/v1/admin/count', fetcher)

    if (isLoading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>

    if (data.count === 0) {
        return <SetupForm />
    }

    return <div>
        <Outlet />
    </div>
}