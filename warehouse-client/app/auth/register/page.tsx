'use client'

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
    Card,
    CardAction,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Label } from "@/components/ui/label";
import Link from "next/link";
import { useState } from "react";
import useRegister from "@/hooks/useRegister";
import { useRouter } from "next/navigation";


export default function Register() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [email, setEmail] = useState("")
    const [fullname, setFullname] = useState("")

    const { register, loading, error } = useRegister()

    const router = useRouter()

    const handleSubmit = async (e: React.SubmitEvent) => {
        e.preventDefault()

        const token = await register(username, password, email, fullname)

        if (token) {
            localStorage.setItem("TOKEN", token)
            router.replace("/auth/login")
        }
    }


    return (
        <div className="w-full min-h-screen flex bg-black">
            <div className="w-100 rounded-xl mx-auto my-auto">
                <Card className="w-full max-w-sm">
                    <CardHeader>
                        <CardTitle>Create new account</CardTitle>
                        <CardDescription>
                            Enter your email below to login to your account
                        </CardDescription>
                        <CardAction>
                            <Link href={"/auth/login"}>
                                <Button variant="link" className="cursor-pointer">Sign In</Button>
                            </Link>
                        </CardAction>
                    </CardHeader>
                    <form onSubmit={handleSubmit}>
                        <CardContent>
                            <div className="flex flex-col gap-6">
                                <div className="grid gap-2">
                                    <Label htmlFor="username">Username</Label>
                                    <Input
                                        id="username"
                                        type="username"
                                        placeholder="your name"
                                        value={username}
                                        onChange={e => setUsername(e.target.value)}
                                        required
                                    />
                                </div>
                                <div className="grid gap-2">
                                    <div className="flex items-center">
                                        <Label htmlFor="password">Password</Label>
                                    </div>
                                    <Input
                                        id="password"
                                        type="password"
                                        value={password}
                                        onChange={e => setPassword(e.target.value)}
                                        required
                                    />
                                </div>
                                <div className="grid gap-2">
                                    <Label htmlFor="email">Email</Label>
                                    <Input
                                        id="email"
                                        type="email"
                                        placeholder="m@example.com"
                                        value={email}
                                        onChange={e => setEmail(e.target.value)}
                                        required
                                    />
                                </div>
                                <div className="grid gap-2">
                                    <Label htmlFor="email">Fullname</Label>
                                    <Input
                                        id="fullname"
                                        type="fullname"
                                        placeholder="fullname"
                                        value={fullname}
                                        onChange={e => setFullname(e.target.value)}
                                        required
                                    />
                                </div>
                            </div>
                            {error && <p className="text-red-500">{error}</p>}
                        </CardContent>
                        <CardFooter className="flex-col gap-2 mt-4">
                            <Button type="submit" className="w-full cursor-pointer">
                                {loading ? "Creating an account..." : "Create new account"}
                            </Button>
                        </CardFooter>
                    </form>
                </Card>
            </div>
        </div >
    );
}

