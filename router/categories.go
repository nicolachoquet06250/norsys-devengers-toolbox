package router

import (
	_ "embed"
	"encoding/base64"
	"norsys-devengers-toolbox/api"
)

//go:embed templates/logos/typescript.png
var typescriptLogo []byte

//go:embed templates/logos/logomui.png
var muiLogo []byte

//go:embed templates/logos/vue-macros.svg.txt
var logoVueMacros string

//go:embed templates/logos/svelte-use.png.txt
var logoSvelteUse string

//go:embed templates/logos/vue-use.svg.txt
var logoVueUse string

//go:embed templates/logos/vuetify.svg.txt
var logoVuetify string

//go:embed templates/logos/gradiant-generator.png.txt
var logoGradiantGenerator string

//go:embed templates/logos/10015.ico.txt
var logo10015 string

//go:embed templates/logos/readme.svg.txt
var logoReadme string

//go:embed templates/logos/omatsuri.svg.txt
var logoOmatsuri string

//go:embed templates/logos/base64-image-encoder.png.txt
var logoBase64ImageEncoder string

//go:embed templates/logos/colorpicker.webp.txt
var logoColorPicker string

//go:embed templates/logos/ray.png.txt
var logoRay string

//go:embed templates/logos/nettools.png.txt
var logoNettools string

//go:embed templates/logos/diffchecker.ico.txt
var logoDiffchecker string

//go:embed templates/logos/imagetotext.svg.txt
var logoImageToText string

//go:embed templates/logos/getnada.ico.txt
var logoNada string

//go:embed templates/logos/regex101.png.txt
var logoRegex101 string

var categories = api.Categories{
	{
		Name: "Front",
		Tools: api.Tools{
			{
				Id:          0,
				Name:        "Vue Macros",
				Description: "Explore and extend more macros and syntax sugar to Vue.",
				Url:         "https://vue-macros.sxzz.moe/",
				Logo:        logoVueMacros,
				Editor:      "Kevin Deng",
				Categorie:   "Front",
			},
			{
				Id:          1,
				Name:        "Svelte Use",
				Description: "Ensemble de hooks (use...) prédévelopés dans svelte.js.",
				Url:         "https://svelte-use.vercel.app/",
				Logo:        logoSvelteUse,
				Editor:      "2nthony",
				Categorie:   "Front",
			},
			{
				Id:          2,
				Name:        "Vue Use",
				Description: "Ensemble de hooks (use...) prédévelopés dans vue.js.",
				Url:         "https://vueuse.org/",
				Logo:        logoVueUse,
				Editor:      "Anthony Fu",
				Categorie:   "Front",
			},
			{
				Id:          3,
				Name:        "React Use",
				Description: "Ensemble de hooks (use...) prédévelopés dans react.js.",
				Url:         "https://github.com/streamich/react-use",
				Logo:        "data:image/png;base64," + base64.StdEncoding.EncodeToString(typescriptLogo),
				Editor:      "streamich",
				Categorie:   "Front",
			},
			{
				Id:          4,
				Name:        "MUI",
				Description: "MUI offers a comprehensive suite of UI tools to help you ship new features faster. Start with Material UI, our fully-loaded component library, or bring your own design system to our production-ready components.",
				Url:         "https://mui.com/",
				Logo:        "data:image/png;base64," + base64.StdEncoding.EncodeToString(muiLogo),
				Editor:      "Material UI SAS",
				Categorie:   "Front",
			},
			{
				Id:          12,
				Name:        "Vuetify",
				Description: "Vuetify is a no design skills required UI Library with beautifully handcrafted Vue Components.",
				Url:         "https://vuetifyjs.com/en/",
				Logo:        logoVuetify,
				Editor:      "Vuetify",
				Categorie:   "Front",
			},
			{
				Id:          14,
				Name:        "Gradient Generator",
				Description: "Beautiful, lush gradients ✨",
				Url:         "https://www.joshwcomeau.com/gradient-generator/",
				Logo:        logoGradiantGenerator,
				Editor:      "Josh Comeau",
				Categorie:   "Front",
			},
			{
				Id:          15,
				Name:        "10015",
				Description: "All Online Tools in “One Box”",
				Url:         "https://10015.io/",
				Logo:        logo10015,
				Editor:      "10015",
				Categorie:   "Front",
			},
		},
	},
	{
		Name: "Edition de fichier",
		Tools: api.Tools{
			{
				Id:          5,
				Name:        "Readme.so",
				Description: "Éditeur visuel de fichiers readme. Permet d'accélérer la création de readme complexes.",
				Url:         "https://readme.so/fr",
				Logo:        logoReadme,
				Editor:      "Katherine Oelsner",
				Categorie:   "Edition de fichier",
			},
		},
	},
	{
		Name: "Docs",
		Tools: api.Tools{
			{
				Id:          6,
				Name:        "Typescript",
				Description: "Documentation de typescript",
				Url:         "https://typescript.org",
				Logo:        "data:image/png;base64," + base64.StdEncoding.EncodeToString(typescriptLogo),
				Editor:      "Microsoft",
				Categorie:   "Docs",
			},
		},
	},
	{
		Name: "Tools",
		Tools: api.Tools{
			{
				Id:          7,
				Name:        "Regex101",
				Description: "Outil Permetant de tester ses regex",
				Url:         "https://regex101.com",
				Logo:        logoRegex101,
				Editor:      "regex101",
				Categorie:   "Tools",
			},
			{
				Id:          8,
				Name:        "Omatsuri",
				Description: "Omatsuri is a progressive web application with 12 open source frontend focused tools. Omatsuri translates to «festival» from Japanese (お祭り) and here we have a small festival of applications. It was built with strong respect to your privacy – you will never see ads and it does not include analytics services (or actually any services at all). You are highly encouraged to explore source code and use it in your projects.",
				Url:         "https://omatsuri.app",
				Logo:        logoOmatsuri,
				Editor:      "OpenSource",
				Categorie:   "Tools",
			},
			{
				Id:          9,
				Name:        "Base64 image encoder",
				Description: "Outil permettant de convertir une image en format url base64",
				Url:         "https://elmah.io/tools/base64-image-encoder/",
				Logo:        logoBase64ImageEncoder,
				Editor:      "elmah.io",
				Categorie:   "Tools",
			},
			{
				Id:          10,
				Name:        "Image Color Picker",
				Description: "Outil permettant générer une palette de couleurs basée sur les couleurs présentent dans une image uploadée.",
				Url:         "https://imagecolorpicker.com/",
				Logo:        logoColorPicker,
				Categorie:   "Tools",
			},
			{
				Id:          11,
				Name:        "Create Beautiful images of your code",
				Description: "Create Beautiful images of your code",
				Url:         "https://ray.so/",
				Logo:        logoRay,
				Categorie:   "Tools",
			},
			{
				Id:          13,
				Name:        "NetTools.club",
				Description: "A fast and easy Online Chmod Calculator to convert Linux file permissions between different formats.",
				Url:         "https://www.nettools.club/chmod_calc",
				Logo:        logoNettools,
				Editor:      "NetTools.club",
				Categorie:   "Tools",
			},
			{
				Id:          16,
				Name:        "DiffChecker",
				Description: "Diffchecker will compare text to find the difference between two text files. Just paste your files and click Find Difference",
				Url:         "https://www.diffchecker.com/text-compare/",
				Logo:        logoDiffchecker,
				Editor:      "Checker Software Inc.",
				Categorie:   "Tools",
			},
			{
				Id:          17,
				Name:        "Convertisseur d'images en texte",
				Description: "Nous vous présentons un service en ligne d'OCR (reconnaissance optique de caractères) pour extraire texte d'une image. Téléchargez votre photo dans notre convertisseur d'image en texte, cliquez sur soumettre et obtenez votre fichier texte instantanément.",
				Url:         "https://www.imagetotext.info/fr/image-en-texte",
				Logo:        logoImageToText,
				Editor:      "Imagetotext.info",
				Categorie:   "Tools",
			},
			{
				Id:          18,
				Name:        "Nada",
				Description: "Disposable Temporary Email, Your inbox is sacred. Our inboxes are throwaway.",
				Url:         "https://getnada.com/",
				Logo:        logoNada,
				Editor:      "Getnada",
				Categorie:   "Tools",
			},
		},
	},
}
