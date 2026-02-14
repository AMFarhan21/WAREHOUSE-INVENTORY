'use client'
import useGetAllBarang from '@/hooks/useGetAllBarang'
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Button } from '@/components/ui/button'
import { Edit, Trash } from 'lucide-react'
import {
    Pagination,
    PaginationContent,
    PaginationEllipsis,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
} from "@/components/ui/pagination"
import useGetAllHistoryStok from '@/hooks/useGetAllHistoryStok'

const Page = () => {
    const { historyStoks, loading, error, meta, setPageNum, pageNum, limitNum } = useGetAllHistoryStok()
    const totalPage = Math.ceil((meta?.total || 0) / limitNum)
    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>History</div>
                <Table className='bg-white rounded-lg w-full mt-4'>
                    {/* <TableCaption>A list of your recent invoices.</TableCaption> */}
                    <TableHeader>
                        <TableRow>
                            <TableHead>Nama Barang</TableHead>
                            <TableHead>Kode Barang</TableHead>
                            <TableHead>Unit</TableHead>
                            <TableHead className="w-[100px]">Keterangan</TableHead>
                            <TableHead>Nama Staff</TableHead>
                            <TableHead className="text-right">Harga Jual</TableHead>
                            <TableHead className="text-right">Action</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {historyStoks.map((historyStok) => (
                            <TableRow key={historyStok.id}>
                                <TableCell className="font-medium">{historyStok.barang.nama_barang}</TableCell>
                                <TableCell>{historyStok.barang.kode_barang}</TableCell>
                                <TableCell>{historyStok.barang.satuan}</TableCell>
                                <TableCell>{historyStok.keterangan}</TableCell>
                                <TableCell>{historyStok.user.full_name}</TableCell>
                                <TableCell className="text-right">Rp. {historyStok.barang.harga_jual}</TableCell>
                                <TableCell className="text-right space-x-4">{historyStok.created_at.split("T")[0] + " | " + historyStok.created_at.split("T")[1].slice(0, 8)}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                    {/* <TableFooter>
                        <TableRow>
                            <TableCell colSpan={3}>Total</TableCell>
                            <TableCell className="text-right">$2</TableCell>
                        </TableRow>
                    </TableFooter> */}

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
        </div>
    )
}

export default Page




