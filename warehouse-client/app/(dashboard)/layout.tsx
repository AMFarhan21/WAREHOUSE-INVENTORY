import { AppSidebar } from "@/components/ui/app-sidebar"
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { Toaster } from "react-hot-toast"

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <SidebarProvider>
            <AppSidebar />
            <div><Toaster /></div>
            <main className="bg-gray-100 w-full">
                <SidebarTrigger />
                {children}
            </main>
        </SidebarProvider>
    )
}