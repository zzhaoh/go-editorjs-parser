# git | gitlab automatic releases

## GitLab: Automatic releases with CI/CD Pipelines

GitLab allows it, to execute commands after repository commits. With that, it’s possible to do automated software tests for code in the repository.

A few weeks ago, a private project required automatic releases on GitLab at given dates – for example monthly. At the moment such a feature is not (yet) supported by GitLab (https://gitlab.com/gitlab-org/gitlab-ce/issues/63858)

However, it is possible to create releases via the GitLab API. So it is just necessary, to use this API, to allow a release in the CI/CD Process. For creating releases via this API there’s already a working Docker-Image which we will use. (https://github.com/inetprocess/gitlab-release)

So what do we need to do, to get things done? How can we schedule automatic releases?

A release can only be created for Git-Tags. So the first step is to automatically create such a Tag. For this, we add a section to our .gitlab-ci.yml File, which is only executed, if it was called by the scheduler. (Which we will set up later)

```yaml
on-schedule:
  stage: build
  only:
    - schedules
  script:
    - git config user.email "${GITLAB_USER_EMAIL}"
    - git config user.name "${GITLAB_USER_NAME}"
    - git remote add api-origin https://oauth2:${GITLAB_ACCESS_TOKEN}@gitlab.com/${CI_PROJECT_PATH}
    - git tag -a "Release_$(date +%Y-%m-%d)" -m "Auto-Release"
    - git push api-origin "Release_$(date +%Y-%m-%d)"
```

In the example above, we use the API with an Access-Token authentification. To do so, such an Access-Token needs to be created. In the user settings under Access Tokens you can create such a token. The required permissions for this use case are api and write_repository.

![](https://www.philipp-doblhofer.at/wp-content/uploads/2019/11/create_gitlab_api_key-1024x588.jpg)

After a click on “Create personal access token”, the token is shown only once. Just copy this value and store it as a variable in the Project settings CI / CD in the Section Variables. For our code example we need to call the variable GITLAB_ACCESS_TOKEN:

![](https://www.philipp-doblhofer.at/wp-content/uploads/2019/11/gitlab-auth-token-variable-1024x539.jpg)

Important – especially for public repositories – is, to enable the Protected and Masked switches. Otherwise, other users can see and use your access token.
Because Project-Maintainers can also read protected variables, we created an extra, restricted user just for doing the release job.

To access the variable in the automatically created Tag, we need to protect this tag (otherwise protected variables can not be used). In our example above we defined that newly created Tags start with “Release”. So we will protect every Tag, which starts with that pattern:

![](https://www.philipp-doblhofer.at/wp-content/uploads/2019/11/gitlab_protect_tag-1024x423.jpg)

After this step is completed, we can now configure the scheduler:

![](https://www.philipp-doblhofer.at/wp-content/uploads/2019/11/gitlab_scheduler-1024x503.jpg)

Here you can configure, how often the scheduler will be executed – so how often our release should be created. If none of the given settings is appropriate, we can define our own interval as a [CRON-Format](https://en.wikipedia.org/wiki/Cron#Overview). Since we defined in .gitlab-ci.yml that if it was called by the scheduler, a Tag should be created, this step is also completed.

To create a release from this Tag, we have to add another section to our .gitlab-ci.yml File, which is only executed, if it’s a Tag:

```yaml
publish:
    image: inetprocess/gitlab-release
    stage: publish
    only:
        - tags
    except:
        - schedules
    dependencies: 
        - build
    script:
        - gitlab-release --message 'Automatic release' ./Demo.txt
```

After that, every time the scheduler is executed, a Tag will be created from the Master-Branch. This Tag will then be released automatically.

To wrap everything up, here a complete demo .gitlab-ci.yml-File:

```yaml
stages:
    - build
    - publish
    
build:
  stage: build
  except:
    - schedules
  script:
    # Usually the application is compiled here
    - echo "Demo" > Demo.txt
  artifacts:
    paths:
      - "./*"

on-schedule:
  stage: build
  image: 
    name: alpine/git
    entrypoint: [""]
  only:
    - schedules
  script:
    - git config user.email "${GITLAB_USER_EMAIL}"
    - git config user.name "${GITLAB_USER_NAME}"
    - git remote add demo-tag-origin https://oauth2:${GITLAB_ACCESS_TOKEN}@gitlab.com/${CI_PROJECT_PATH}
    - git tag -a "Release_$(date +%Y-%m-%d)" -m "Auto-Release"
    - git push demo-tag-origin "Release_$(date +%Y-%m-%d)"
    
publish:
    image: inetprocess/gitlab-release
    stage: publish
    only:
        - tags
    except:
        - schedules
    dependencies: 
        - build
    script:
        - gitlab-release --message 'Automatic release' ./Demo.txt
```