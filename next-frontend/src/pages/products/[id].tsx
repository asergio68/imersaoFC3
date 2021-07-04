// @flow
import { GetServerSideProps, NextPage } from "next";
import http from "../../http";
import { Product } from "../../model";

type ProductDetailProps = {
  product: Product;
};
const ProductDetailPage: NextPage<ProductDetailProps> = ({
  product,
  children,
}) => {
  return (
    <div>
      <h3>Produto - {product.name}</h3>
      <ul>
        <li>
          {product.id} | {product.price}
        </li>
      </ul>
    </div>
  );
};

export default ProductDetailPage;

export const getServerSideProps: GetServerSideProps<
  ProductDetailProps,
  { id: string }
> = async (context) => {
  const { id } = context.params!;
  const { data: product } = await http.get(`products/${id}`);
  console.log(product);
  return {
    props: { product },
  };
};
