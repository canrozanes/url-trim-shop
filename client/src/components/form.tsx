import React, { useState, MouseEvent } from "react"
import { styled } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Link from "@material-ui/core/Link"
import Button from "@material-ui/core/Button"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"
import CircularProgress from "@material-ui/core/CircularProgress"
import postData from "../services/index"

const StyledButton = styled(Button)({
  marginLeft: "20px",
  padding: "16px 0",
  width: "calc(100% - 20px)",
})

const SmallButton = styled(Button)({
  marginLeft: "20px",
  width: "100px",
})

const StyledGrid = styled(Grid)(({ theme }) => {
  return {
    borderRadius: "4px",
    border: `2px solid ${theme.palette.primary.main}`,
    padding: "16px 20px",
    marginTop: "20px",
    backgroundColor: theme.palette.primary.light,
  }
})

const StyledTextField = styled(TextField)({
  maxWidth: "100%",
})

const isNotEmpty = (string: string) =>
  typeof string === "string" && string.length > 0

const Form = () => {
  const [urlField, setUrlField] = useState("")
  const [targetUrlField, setTargetUrlField] = useState("")
  const [shortenedUrl, setShortenedUrl] = useState("")
  const [isFetching, setIsFetching] = useState(false)
  const [isClicked, setIsClicked] = useState(false)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrlField(event.target.value)
  }
  const handleClick = async (e: MouseEvent<HTMLButtonElement>) => {
    setIsFetching(true)
    setTargetUrlField(urlField)
    e.preventDefault()
    try {
      const res = await postData(targetUrlField)
      setShortenedUrl(`${window.location}${res.body.Hash}`)
      setIsFetching(false)
    } catch (e) {
      console.log(e)
      setIsFetching(false)
    }
  }

  const copyToClipboard = (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault()
    setIsClicked(true)
    const tempInput = document.createElement("input")
    tempInput.value = shortenedUrl
    document.body.appendChild(tempInput)
    tempInput.select()
    document.execCommand("copy")
    document.body.removeChild(tempInput)
    setTimeout(() => {
      setIsClicked(false)
    }, 700)
  }

  return (
    <>
      <Grid container component="form" alignItems="center">
        <Grid item xs={12} md={9}>
          <StyledTextField
            label="Shorten your link"
            variant="filled"
            onChange={handleChange}
            fullWidth
          />
        </Grid>
        <Grid item xs={12} md={3}>
          <StyledButton
            type="submit"
            id="submitButton"
            onClick={handleClick}
            variant="contained"
            color="primary"
            disabled={isFetching}
            endIcon={
              isFetching ? (
                <CircularProgress
                  color="secondary"
                  size={20}
                  aria-describedby="submitButton"
                  aria-busy={true}
                />
              ) : null
            }
          >
            Shorten
          </StyledButton>
        </Grid>
      </Grid>
      {isNotEmpty(shortenedUrl) && (
        <StyledGrid container alignItems="center">
          <Grid item xs={12} md={6}>
            <Typography variant="body1" component="span">
              Your original URL: {`http://${targetUrlField}`}
            </Typography>
          </Grid>
          <Grid item xs={12} md={6}>
            <Grid container justify="flex-end" alignItems="center">
              <Grid item>
                <Typography variant="body1" component="span" align="right">
                  Your shortened URL:{" "}
                  <Link href={shortenedUrl}>{shortenedUrl}</Link>
                </Typography>
              </Grid>
              <Grid item>
                <SmallButton
                  variant="contained"
                  color={isClicked ? "default" : "secondary"}
                  onClick={copyToClipboard}
                >
                  {isClicked ? "Copied!" : "Copy"}
                </SmallButton>
              </Grid>
            </Grid>
          </Grid>
        </StyledGrid>
      )}
    </>
  )
}

export default Form
