import React from "react";
import { useGetRecipesQuery } from "../generated/graphql";
import styled from "styled-components";
import { useTable, Column, CellProps } from "react-table";
import { Link } from "react-router-dom";
import Debug from "../components/Debug";

interface TableProps<T extends object> {
  columns: Column<T>[];
  data: T[];
}

const Table = <T extends object>({ columns, data }: TableProps<T>) => {
  // Use the state and functions returned from useTable to build your UI
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
  } = useTable({
    columns,
    data,
  });

  // Render the UI for your table
  return (
    <table {...getTableProps()} data-cy="recipe-table">
      <thead>
        {headerGroups.map((headerGroup) => (
          <tr {...headerGroup.getHeaderGroupProps()}>
            {headerGroup.headers.map((column) => (
              <th {...column.getHeaderProps()}>{column.render("Header")}</th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody {...getTableBodyProps()}>
        {rows.map((row) => {
          prepareRow(row);
          return (
            <tr {...row.getRowProps()}>
              {row.cells.map((cell) => {
                return <td {...cell.getCellProps()}>{cell.render("Cell")}</td>;
              })}
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

const Styles = styled.div`
  padding: 1rem;

  table {
    border-spacing: 0;
    border: 1px solid black;

    tr {
      :last-child {
        td {
          border-bottom: 0;
        }
      }
    }

    th,
    td {
      margin: 0;
      padding: 0.5rem;
      border-bottom: 1px solid black;
      border-right: 1px solid black;

      :last-child {
        border-right: 0;
      }
    }
  }
`;

const RecipeList: React.FC = () => {
  const { data, error } = useGetRecipesQuery({});

  const columns = React.useMemo(
    () => [
      {
        Header: "UUID",
        accessor: "uuid",
      },
      {
        Header: "Name",
        accessor: "name",
      },
      {
        Header: "test",
        accessor: "test",
        Cell: (cell: CellProps<any>) => (
          <Link to={`recipe/${cell.row.original.uuid}`} className="link">
            details
          </Link>
        ),
      },
    ],
    []
  );

  return (
    <Styles>
      <Table columns={columns} data={data?.recipes || []} />
      <Debug data={{ error }} />
    </Styles>
  );
};

export default RecipeList;
