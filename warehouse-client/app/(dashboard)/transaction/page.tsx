'use client'
import { Button } from '@/components/ui/button'
import { Coins } from 'lucide-react'
import Link from 'next/link'

const Page = () => {
    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>TRANSACTION</div>
                <div className='flex gap-8 mt-12'>
                    <Link href={"/transaction/pembelian"}>
                        <Button className='text-4xl w-100 h-50 bg-blue-200 text-black hover:bg-blue-300'>
                            Pembelian
                        </Button>
                    </Link>
                    <Link href={"/transaction/penjualan"}>
                        <Button className='text-4xl w-100 h-50 bg-green-200 text-black hover:bg-green-300'>
                            Penjualan
                        </Button>
                    </Link>
                </div>

            </div>
        </div>
    )
}

export default Page