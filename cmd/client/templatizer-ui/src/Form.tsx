import { Button, Grid, Paper, Typography } from "@material-ui/core";
import axios, { AxiosResponse } from "axios";
import { useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { RepositoryDetails, defaultValues as repoDetailsDefaultValues } from './components/RepositoryDetails';
import { TemplateVars } from './components/TemplateVars';
import { Config, LoadParameters, ProcessTemplate, ProcessTemplateRequest, LoadParametersRequest } from './components/Client';


interface IFormInput {
  url: string;
  authMechanism: string;
  authUsername: string;
  authPassword: string;
  authToken: string;
  branchDefault: string;
  branch: string;
  configPath: string;
  templateVars: Config;
  params: {}
  [index: string]: any;
  mode: string;
}


const defaultValues = {
  ...repoDetailsDefaultValues,
  params: {},
};

export const TemplatizerForm = () => {

  const methods = useForm<IFormInput>({ defaultValues: defaultValues });
  const { control, setValue, getValues, formState } = methods;
  const [showTemplateVars, setTemplateVars] = useState<boolean>(false);

  const buildLoadParametersRequest = (): LoadParametersRequest => {
    let request: LoadParametersRequest = {
      url: getValues('url'),
    }
    let authMechanism = getValues('authMechanism')
    if (authMechanism == 'basic') {
      request.auth = {
        mechanism: authMechanism,
        username: getValues('authUsername'),
        password: getValues('authPassword'),
      }
    } else if (authMechanism == 'token') {
      request.auth = {
        mechanism: authMechanism,
        token: getValues('authToken'),
      }
    }
    if (getValues('branchDefault') != 'default') {
      request.branch = getValues('branch')
    }
    return request
  }

  const loadVariables = () => {
    const request = buildLoadParametersRequest()
    LoadParameters(request).then((response: AxiosResponse<Config>) => {
      setValue("templateVars", response.data)
      setTemplateVars(true)
    }).catch((error) => {
      console.log(error.response);
      alert(error.response.data.message);
    });
  }

  const processTemplate = () => {
    let params: any = getValues('params')
    const request: ProcessTemplateRequest = {
      ...buildLoadParametersRequest(),
      params,
    }

    ProcessTemplate(request).then((response: AxiosResponse) => {
      console.log(response.headers)
      let fileName = response.headers["content-disposition"].split("filename=")[1];
      const url = window.URL.createObjectURL(new Blob([response.data],
        { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' }));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download',
        response.headers["content-disposition"].split("filename=")[1]);
      document.body.appendChild(link);
      link.click();

    }).catch((err: any) => {
      console.log(err)
      alert(err.data);
    });
  }

  const title = "{{.Templatizer}}"

  const updateParam = (name: string, value: any) => {
    let params: any = getValues('params')
    params[name] = value
    setValue('params', params)
  }

  return (
    <Paper
      style={{
        display: "grid",
        gridRowGap: "20px",
        padding: "20px",
        margin: "10px 300px",
      }}
    >
      <Typography variant="h3">{title}</Typography>
      {!showTemplateVars &&
        <Grid container spacing={1}>
          <RepositoryDetails control={control} setValue={setValue} />
          <hr />
          <Button className="btn-primary" onClick={() => loadVariables()} variant={"outlined"}>
            {" "}Next{" "}
          </Button>
        </Grid>
      }
      {showTemplateVars &&
        <Grid container spacing={1}>

          <Grid item lg={12}>
            <Typography variant="h6">Tempalte variables</Typography>
          </Grid>
          <TemplateVars control={control} setValue={setValue} updateParam={updateParam} />
          <Button onClick={() => {
            setTemplateVars(false)
          }} variant={"outlined"}>
            {" "}Back{" "}
          </Button>
          <hr />
          <Button onClick={() => processTemplate()} variant={"outlined"}>
            {" "}
            Process Template{" "}
          </Button>
        </Grid>
      }

    </Paper>
  );
};
