'use client'
import { ShoppingBag, ShoppingCart } from 'lucide-react'
import Link from 'next/link'

import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "@/components/ui/accordion"
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import {
    Pagination,
    PaginationContent,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
} from "@/components/ui/pagination"
import useGetAllPembelian from '@/hooks/useGetAllPembelian'
import useGetAllPenjualan from '@/hooks/useGetAllPenjualan'

const Page = () => {

    const { pembelians, meta, limitNum, pageNum, setPageNum } = useGetAllPembelian()
    const { penjualans, metaPenjualan, limitNumPenjualan, pageNumPenjualan, setPageNumPenjualan } = useGetAllPenjualan()
    const totalPage = Math.ceil((meta?.total || 0) / limitNum)
    const totalPagePenjualan = Math.ceil((metaPenjualan?.total || 0) / limitNum)

    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>TRANSACTION</div>
                <div className='flex gap-4 mt-12'>
                    <Link href="/transaction/pembelian">
                        <button className='bg-blue-200 font-bold gap-2 shadow-xl text-4xl rounded-xl cursor-pointer text-black hover:bg-blue-300 w-64 h-32 flex items-center justify-center'>
                            <span>Beli</span>
                            <ShoppingCart className='w-8 h-8' />
                        </button>
                    </Link>
                    <Link href="/transaction/penjualan">
                        <button className='bg-green-200 font-bold gap-2 shadow-xl text-4xl rounded-xl cursor-pointer  text-black hover:bg-green-300 w-64 h-32 flex items-center justify-center'>
                            <span>Jual</span>
                            <ShoppingBag className='w-8 h-8' />
                        </button>
                    </Link>
                </div>

                <div className="w-full grid grid-cols-2 gap-4">
                    <div>
                        <div className='text-xl font-bold mt-12'>Pembelian</div>
                        <Table className='bg-white rounded-lg w-full mt-4'>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Nomor Faktur</TableHead>
                                    <TableHead>Supplier</TableHead>
                                    <TableHead>Total</TableHead>
                                    <TableHead>Status</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {pembelians && pembelians.map((pembelian) => (
                                    <TableRow key={pembelian.id}>
                                        <TableCell>{pembelian.no_faktur}</TableCell>
                                        <TableCell>{pembelian.supplier}</TableCell>
                                        <TableCell>Rp.{pembelian.total}</TableCell>
                                        <TableCell>{pembelian.status}</TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                        <Pagination className='mt-4'>
                            <PaginationContent>
                                <PaginationItem>
                                    <PaginationPrevious className='cursor-pointer' onClick={() => pageNum > 1 && setPageNum(pageNum - 1)} />
                                </PaginationItem>
                                {
                                    Array.from({ length: totalPage }).map((_, i) => (
                                        <PaginationItem className='cursor-pointer' key={i}>
                                            <PaginationLink onClick={() => {
                                                setPageNum(i + 1)
                                            }} isActive={pageNum == i + 1}>
                                                {i + 1}
                                            </PaginationLink>
                                        </PaginationItem>
                                    ))
                                }
                                <PaginationItem>
                                    <PaginationNext className='cursor-pointer' onClick={() => pageNum != totalPage && setPageNum(pageNum + 1)} />
                                </PaginationItem>
                            </PaginationContent>
                        </Pagination>
                    </div>

                    <div>
                        <div className='text-xl font-bold mt-12'>Penjualan</div>
                        <Table className='bg-white rounded-lg w-full mt-4'>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Nomor Faktur</TableHead>
                                    <TableHead>Supplier</TableHead>
                                    <TableHead>Total</TableHead>
                                    <TableHead>Status</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {penjualans && penjualans.map((penjualan) => (
                                    <TableRow key={penjualan.id}>
                                        <TableCell>{penjualan.no_faktur}</TableCell>
                                        <TableCell>{penjualan.customer}</TableCell>
                                        <TableCell>Rp.{penjualan.total}</TableCell>
                                        <TableCell>{penjualan.status}</TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                        <Pagination className='mt-4'>
                            <PaginationContent>
                                <PaginationItem>
                                    <PaginationPrevious className='cursor-pointer' onClick={() => pageNumPenjualan > 1 && setPageNumPenjualan(pageNum - 1)} />
                                </PaginationItem>
                                {
                                    Array.from({ length: totalPagePenjualan }).map((_, i) => (
                                        <PaginationItem className='cursor-pointer' key={i}>
                                            <PaginationLink onClick={() => {
                                                setPageNumPenjualan(i + 1)
                                            }} isActive={pageNumPenjualan == i + 1}>
                                                {i + 1}
                                            </PaginationLink>
                                        </PaginationItem>
                                    ))
                                }
                                <PaginationItem>
                                    <PaginationNext className='cursor-pointer' onClick={() => pageNumPenjualan != totalPagePenjualan && setPageNumPenjualan(pageNum + 1)} />
                                </PaginationItem>
                            </PaginationContent>
                        </Pagination>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Page