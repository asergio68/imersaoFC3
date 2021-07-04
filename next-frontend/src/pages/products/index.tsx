import { GetServerSideProps, GetStaticProps, NextPage } from "next";
import http from "../../http";
import { Product } from "../../model";
import Link from "next/link";

interface ProductsListPageProps {
  products: Product[];
}
const ProductsListPage: NextPage<ProductsListPageProps> = ({ products }) => {
  return (
    <div>
      <h3>Listagem de produtos</h3>
      <ul>
        {products.map((product, key) => (
          <li key={key}>
            {product.id} | {product.name} |{" "}
            <Link href="/products/[id]" as={`/products/${product.id}`}>
              <a>Detalhes</a>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductsListPage;

export const getServerSideProps: GetServerSideProps<ProductsListPageProps> =
  async (context) => {
    console.log("products");
    const { data: products } = await http.get("products");
    console.log(products);
    return {
      props: {
        products,
      },
    };
  };
