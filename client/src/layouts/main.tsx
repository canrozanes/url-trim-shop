import React from "react"
import { Helmet } from "react-helmet"
import styled from "styled-components"
import Container from "@material-ui/core/Container"
import CssBaseline from "@material-ui/core/CssBaseline"
import Footer from "../components/footer"
import Header from "../components/header"

const PageContainer = styled("div")({
  minHeight: "100vh",
  display: "flex",
  flexDirection: "column",
})

const MainContent = styled(Container)({
  flex: 1,
})

interface LayoutProps {
  children: JSX.Element[] | JSX.Element
}

const MainLayout = ({ children }: LayoutProps) => (
  <>
    <Helmet>
      <meta charSet="utf-8" />
      <meta
        name="viewport"
        content="minimum-scale=1, initial-scale=1, width=device-width"
      />
      <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
      />
      <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/icon?family=Material+Icons"
      />
      <title>url-trim.shop</title>
    </Helmet>
    <PageContainer>
      <Header />
      <CssBaseline />
      <MainContent>{children}</MainContent>
      <Footer />
    </PageContainer>
  </>
)

export default MainLayout
