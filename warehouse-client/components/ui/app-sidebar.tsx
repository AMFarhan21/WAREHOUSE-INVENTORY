'use client'
import * as React from "react"

import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarRail,
} from "@/components/ui/sidebar"
import { VersionSwitcher } from "./version.switcher"
import { usePathname } from "next/navigation"
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import { Button } from "@/components/ui/button"
import Link from "next/link"

// This is sample data.
const data = {
    versions: ["1.0.1", "1.1.0-alpha", "2.0.0-beta1"],
    navMain: [
        // {
        //     title: "Getting Started",
        //     url: "#",
        //     items: [
        //         {
        //             title: "Installation",
        //             url: "#",
        //         },
        //         {
        //             title: "Project Structure",
        //             url: "#",
        //         },
        //     ],
        // },
        {
            title: "Menu",
            url: "#",
            items: [
                {
                    title: "Dashboard",
                    url: "/dashboard",
                    // isActive: true,
                },
                {
                    title: "Transaction",
                    url: "/transaction",
                    // isActive: true,
                },
                {
                    title: "History",
                    url: "/history",
                    // isActive: true,
                },
                {
                    title: "Logout",
                    url: "/auth/login",
                    // isActive: true,
                },
                // {
                //     title: "Styling",
                //     url: "#",
                // },
                // {
                //     title: "Optimizing",
                //     url: "#",
                // },
                // {
                //     title: "Configuring",
                //     url: "#",
                // },
                // {
                //     title: "Testing",
                //     url: "#",
                // },
                // {
                //     title: "Authentication",
                //     url: "#",
                // },
                // {
                //     title: "Deploying",
                //     url: "#",
                // },
                // {
                //     title: "Upgrading",
                //     url: "#",
                // },
                // {
                //     title: "Examples",
                //     url: "#",
                // },
            ],
        },
        // {
        //     title: "API Reference",
        //     url: "#",
        //     items: [
        //         {
        //             title: "Components",
        //             url: "#",
        //         },
        //         {
        //             title: "File Conventions",
        //             url: "#",
        //         },
        //         {
        //             title: "Functions",
        //             url: "#",
        //         },
        //         {
        //             title: "next.config.js Options",
        //             url: "#",
        //         },
        //         {
        //             title: "CLI",
        //             url: "#",
        //         },
        //         {
        //             title: "Edge Runtime",
        //             url: "#",
        //         },
        //     ],
        // },
        // {
        //     title: "Architecture",
        //     url: "#",
        //     items: [
        //         {
        //             title: "Accessibility",
        //             url: "#",
        //         },
        //         {
        //             title: "Fast Refresh",
        //             url: "#",
        //         },
        //         {
        //             title: "Next.js Compiler",
        //             url: "#",
        //         },
        //         {
        //             title: "Supported Browsers",
        //             url: "#",
        //         },
        //         {
        //             title: "Turbopack",
        //             url: "#",
        //         },
        //     ],
        // },
    ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
    const pathname = usePathname()
    return (
        <Sidebar {...props}>
            <SidebarHeader>
                <VersionSwitcher
                    versions={data.versions}
                    defaultVersion={data.versions[0]}
                />
                {/* <SearchForm /> */}
            </SidebarHeader>
            <SidebarContent>
                {/* We create a SidebarGroup for each parent. */}
                {data.navMain.map((nav) => (
                    <SidebarGroup key={nav.title}>
                        <SidebarGroupLabel>{nav.title}</SidebarGroupLabel>
                        <SidebarGroupContent>
                            <SidebarMenu>
                                {nav.items.map((item, i) => (
                                    <SidebarMenuItem key={item.title}>
                                        {
                                            item.title == "Logout" ? (
                                                <SidebarMenuButton asChild>
                                                    <AlertDialog>
                                                        <AlertDialogTrigger asChild>
                                                            <Button className=" cursor-pointer mt-2">{item.title}</Button>
                                                        </AlertDialogTrigger >

                                                        <AlertDialogContent>
                                                            <AlertDialogHeader>
                                                                <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                                                                <AlertDialogDescription>
                                                                    This action cannot be undone. This will log you out
                                                                </AlertDialogDescription>
                                                            </AlertDialogHeader>
                                                            <AlertDialogFooter>
                                                                <AlertDialogCancel className="cursor-pointer">Cancel</AlertDialogCancel>
                                                                <Link href={"auth/login"}>
                                                                    <AlertDialogAction className="bg-red-500 cursor-pointer" onClick={() => {
                                                                        localStorage.removeItem("TOKEN")
                                                                    }}>Continue</AlertDialogAction>
                                                                </Link>
                                                            </AlertDialogFooter>
                                                        </AlertDialogContent>
                                                    </AlertDialog>
                                                </SidebarMenuButton>


                                            ) : (
                                                <SidebarMenuButton asChild isActive={pathname == item.url}>
                                                    <a href={item.url}>{item.title}</a>
                                                </SidebarMenuButton>
                                            )
                                        }
                                    </SidebarMenuItem>
                                ))}
                            </SidebarMenu>
                        </SidebarGroupContent>
                    </SidebarGroup>
                ))}
            </SidebarContent>
            <SidebarRail />
        </Sidebar>
    )
}