// var Generator = require('yeoman-generator');

// module.exports = class extends Generator {
//   method1() {
//     this.log('method 1 just ran');
//   }

//   method2() {
//     this.log('method 2 just ran');
//   }
// };

'use strict';

var util       = require('util'),
    path       = require('path'),
    generators = require('yeoman-generator'),
    _          = require('lodash'),
    _s         = require('underscore.string'),
    pluralize  = require('pluralize'),
    asciify    = require('asciify');


module.exports = generators.Base.extend({
  constructor: function () {
    generators.Base.apply(this, arguments);

    this.option('flat', {
      type: Boolean,
      required: false,
      defaults: false,
      desc: 'When specified, generators will be created at the top level of the project.'
    });
  },

  initializing: function () {
    this.config.set('structure', this.options.flat ? 'flat' : 'nested');
    this.generatorsPrefix = this.options.flat ? '' : 'generators/';
    this.appGeneratorDir = this.options.flat ? 'app' : 'generators';
  },

  prompting: {
    askFor: function () {
      var done = this.async();
      this.log('\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '|g|o| |l|a|m|b|d|a| |g|e|n|e|r|a|t|o|r|\n' +
        '+-+-+ +-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
        '\n'
      );

      var prompts = [{
        type: 'input',
        name: 'lambdaName',
        message: 'What is the name of your Lambda function?',
        store   : true,
        default: 'mylambda'
      }
    ];

      this.prompt(prompts, function (props) {
        this.lambdaName = props.lambdaName;

        done();
      }.bind(this));
    }
  },

  writing: {
    app: function () {

      var infrastructureDir          = 'infrastructure/',
          terraformDir               = 'infrastructure/terraform/',
          terraformBackendConfigsDir = 'infrastructure/terraform/backendConfigs/',
          handlerDir                 = 'handler/',
          processDir                 = 'process/';

      this.mkdir(infrastructureDir);
      this.mkdir(terraformDir);
      this.mkdir(handlerDir);
      this.mkdir(processDir);
      this.mkdir(terraformBackendConfigsDir);

      this.fs.copyTpl(
        this.templatePath('main.go'),
        this.destinationPath('main.go'),
        {
          lambdaName: this.lambdaName
        }
      );

      this.fs.copyTpl(
        this.templatePath('README.md.erb'),
        this.destinationPath('README.md'),
        {
          lambdaName: this.lambdaName
        }
      );

      this.fs.copyTpl(
        this.templatePath('Makefile'),
        this.destinationPath('Makefile'),
        {
          lambdaName: this.lambdaName
        }
      );

      this.fs.copyTpl(
        this.templatePath('gitignore'),
        this.destinationPath('.gitignore'),
        {
          lambdaName: this.lambdaName
        }
      );

      this.fs.copyTpl(
        this.templatePath('Gopkg.toml'),
        this.destinationPath('Gopkg.toml')
      );

      this.fs.copyTpl(
        this.templatePath('Gopkg.lock'),
        this.destinationPath('Gopkg.lock')
      );

      this.fs.copyTpl(
        this.templatePath('vault-cas.crt'),
        this.destinationPath('vault-cas.crt')
      );
      
      this.template(handlerDir + 'handler.go', handlerDir + 'handler.go');
      this.template(handlerDir + 'initialize.go', handlerDir + 'initialize.go');

      this.template(processDir + 'process.go', processDir + 'process.go');
      this.template(processDir + 'process_test.go', processDir + 'process_test.go');

      this.template(terraformDir + 'provider.tf', terraformDir + 'provider.tf');
      this.template(terraformDir + 'backend.tf', terraformDir + 'backend.tf');
      this.template(terraformDir + 'data.tf', terraformDir + 'data.tf');
      this.template(terraformDir + 'main.tf', terraformDir + 'main.tf');
      this.template(terraformDir + 'outputs.tf', terraformDir + 'outputs.tf');
      this.template(terraformDir + 'variables.tf', terraformDir + 'variables.tf');
      
      this.template(terraformDir + 'setEnv.sh', terraformDir + 'setEnv.sh');

      this.template(terraformBackendConfigsDir + 'dev', terraformBackendConfigsDir + 'dev');
      this.template(terraformBackendConfigsDir + 'qa', terraformBackendConfigsDir + 'qa');
      this.template(terraformBackendConfigsDir + 'prod', terraformBackendConfigsDir + 'prod');
    }
  }
});