import React, {useEffect} from 'react';
import clsx from 'clsx';
import {createStyles, makeStyles, Theme} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import {stringToHex} from "./analytics.utils";

interface Data {
  id: number;
  bank_id: string,
  account_id: string,
  transaction_date: string;
  description: string;
  category: string;
  amount: string;
}

export const mapTransactionToData = t => createData(
  t.TransactionId,
  t.BankId,
  t.AccountId,
  t.BookingDateTime,
  t.TransactionInformation,
  t.BankTransactionCode.Code,
  t.Amount.Amount
);

function createData(
  id: number,
  bank_id: string,
  account_id: string,
  transaction_date: string,
  description: string,
  category: string,
  amount: string,
): Data {
  return {id, bank_id, account_id, transaction_date, description, category, amount};
}

function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

type Order = 'asc' | 'desc';

function getComparator<Key extends keyof any>(
  order: Order,
  orderBy: Key,
): (a: { [key in Key]: number | string }, b: { [key in Key]: number | string }) => number {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort<T>(array: T[], comparator: (a: T, b: T) => number) {
  const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

interface HeadCell {
  disablePadding: boolean;
  id: keyof Data;
  label: string;
  numeric: boolean;
}

const headCells: HeadCell[] = [
  {id: 'transaction_date', numeric: true, disablePadding: false, label: 'Transaction Date'},
  {id: 'account_id', numeric: true, disablePadding: false, label: 'Account information'},
  {id: 'description', numeric: true, disablePadding: false, label: 'Payee'},
  {id: 'category', numeric: true, disablePadding: false, label: 'Category'},
  {id: 'amount', numeric: true, disablePadding: false, label: 'Amount'},
];

interface EnhancedTableProps {
  classes: ReturnType<typeof useStyles>;
  numSelected: number;
  onRequestSort: (event: React.MouseEvent<unknown>, property: keyof Data) => void;
  onSelectAllClick: (event: React.ChangeEvent<HTMLInputElement>) => void;
  order: Order;
  orderBy: string;
  rowCount: number;
}

function EnhancedTableHead(props: EnhancedTableProps) {
  const {classes, order, orderBy, onRequestSort} = props;
  const createSortHandler = (property: keyof Data) => (event: React.MouseEvent<unknown>) => {
    onRequestSort(event, property);
  };

  return (
    <TableHead>
      <TableRow className={'analytics-table-head'}>
        {headCells.map((headCell) => (
          <TableCell
            key={headCell.id}
            align={'left'}
            padding={headCell.disablePadding ? 'none' : 'default'}
            sortDirection={orderBy === headCell.id ? order : false}
          >
            <TableSortLabel
              active={orderBy === headCell.id}
              direction={orderBy === headCell.id ? order : 'asc'}
              onClick={createSortHandler(headCell.id)}
            >
              {headCell.label}
              {orderBy === headCell.id ? (
                <span className={classes.visuallyHidden}>
                  {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                </span>
              ) : null}
            </TableSortLabel>
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
  );
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
    },
    paper: {
      width: '100%',
      height: '100%',
      position: 'relative'
    },
    table: {
      //minWidth: 750,
    },
    visuallyHidden: {
      border: 0,
      clip: 'rect(0 0 0 0)',
      height: 1,
      margin: -1,
      overflow: 'hidden',
      padding: 0,
      position: 'absolute',
      top: 20,
      width: 1,
    },
    tablePagination: {
      position: 'absolute',
      bottom: 0,
      right: 0
    }
  }),
);

type AnalyticsTableProps = {
  data: Data[]
  style: any
}

export default function AnalyticsTable({data, style = {}}: AnalyticsTableProps) {
  const classes = useStyles();
  const [order, setOrder] = React.useState<Order>('asc');
  const [orderBy, setOrderBy] = React.useState<keyof Data>('transaction_date');
  const [selected, setSelected] = React.useState<number[]>([]);
  const [page, setPage] = React.useState(0);
  const [dense] = React.useState(false);
  const [rowsPerPage, setRowsPerPage] = React.useState(5);

  useEffect(() => {
    const rootHeight = document.getElementById('analytics-table-root')?.clientHeight;
    const headHeight = document.getElementsByClassName('analytics-table-head')?.item(0)?.clientHeight;
    const rowHeight = document.getElementsByClassName('analytics-table-row')?.item(0)?.clientHeight;
    const paginationHeight = document.getElementsByClassName('analytics-table-pagination')?.item(0)?.clientHeight;

    if (rootHeight && headHeight && rowHeight && paginationHeight) {
      setRowsPerPage(Math.floor((rootHeight - headHeight - paginationHeight) / rowHeight))
    }
  }, []);

  const handleRequestSort = (event: React.MouseEvent<unknown>, property: keyof Data) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelecteds = data.map((n) => n.id);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const isSelected = (id: number) => selected.indexOf(id) !== -1;

  return (
    <div className={classes.root} id="analytics-table-root" style={style}>
      <Paper className={classes.paper}>
        <TableContainer>
          <Table
            className={classes.table}
            aria-labelledby="tableTitle"
            size={dense ? 'small' : 'medium'}
            aria-label="enhanced table"
          >
            <EnhancedTableHead
              classes={classes}
              numSelected={selected.length}
              order={order}
              orderBy={orderBy}
              onSelectAllClick={handleSelectAllClick}
              onRequestSort={handleRequestSort}
              rowCount={data.length}
            />
            <TableBody>
              {stableSort(data, getComparator(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map((row, index) => {
                  const isItemSelected = isSelected(row.id);
                  const labelId = `enhanced-analytics-table-checkbox-${index}`;

                  return (
                    <TableRow
                      hover
                      className={'analytics-table-row'}
                      // onClick={(event) => handleClick(event, row.id)}
                      role="checkbox"
                      aria-checked={isItemSelected}
                      tabIndex={-1}
                      key={row.id + index}
                      selected={isItemSelected}
                    >
                      <TableCell id={labelId} scope="row" align="left">
                        <span style={{background: '#ECECEC', padding: 4}}>{row.transaction_date}</span>
                      </TableCell>
                      <TableCell align="left">
                        <div style={{display: 'flex', alignItems: 'center', fontSize: 12}}>
                          <Typography variant={'caption'} style={{marginRight: 8}}>Bank ID</Typography>
                          <Typography style={{marginRight: 16}}>{row.bank_id}</Typography>
                          <Typography variant={'caption'} style={{marginRight: 8}}>Account ID</Typography>
                          <Typography>{row.account_id}</Typography>
                        </div>
                        <div>
                          <Typography>**** ***** **** {row.account_id}</Typography>
                        </div>
                      </TableCell>
                      <TableCell align="left">{row.description}</TableCell>
                      <TableCell align="left">
                        <div>
                          <div style={{
                            display: 'inline-block',
                            width: 12,
                            height: 12,
                            borderRadius: '50%',
                            background: `${stringToHex(row.category)}`,
                            position: 'relative',
                            top: 1,
                            marginRight: 4
                          }}/>
                          {row.category}
                        </div>
                      </TableCell>
                      <TableCell align="left">{row.amount}</TableCell>
                    </TableRow>
                  );
                })}
              {data.length === 0 && (
                <TableRow>
                  <TableCell colSpan={6}>No transaction records found</TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          className={clsx(['analytics-table-pagination', classes.tablePagination])}
          rowsPerPageOptions={[]}
          component="div"
          count={data.length}
          rowsPerPage={rowsPerPage}
          page={page}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Paper>
    </div>
  );
}
