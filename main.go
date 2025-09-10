package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Purnachandra - Cloud & DevOps Enthusiast</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
        }
        
        .container {
            max-width: 800px;
            padding: 40px;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 20px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
            border: 1px solid rgba(255, 255, 255, 0.18);
            text-align: center;
        }
        
        .profile-pic {
            width: 150px;
            height: 150px;
            border-radius: 50%%;
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
            margin: 0 auto 30px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 60px;
            font-weight: bold;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
        }
        
        h1 {
            font-size: 3em;
            margin-bottom: 20px;
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text;
        }
        
        .subtitle {
            font-size: 1.5em;
            margin-bottom: 30px;
            opacity: 0.9;
        }
        
        .description {
            font-size: 1.2em;
            line-height: 1.8;
            margin-bottom: 40px;
            opacity: 0.95;
        }
        
        .tech-stack {
            display: flex;
            justify-content: center;
            flex-wrap: wrap;
            gap: 15px;
            margin-bottom: 40px;
        }
        
        .tech-tag {
            background: rgba(255, 255, 255, 0.2);
            padding: 10px 20px;
            border-radius: 25px;
            font-size: 0.9em;
            border: 1px solid rgba(255, 255, 255, 0.3);
            transition: all 0.3s ease;
        }
        
        .tech-tag:hover {
            background: rgba(255, 255, 255, 0.3);
            transform: translateY(-2px);
        }
        
        .connect-btn {
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
            color: white;
            padding: 15px 40px;
            border: none;
            border-radius: 30px;
            font-size: 1.1em;
            font-weight: bold;
            cursor: pointer;
            transition: all 0.3s ease;
            text-decoration: none;
            display: inline-block;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
        }
        
        .connect-btn:hover {
            transform: translateY(-3px);
            box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
        }
        
        .blog-section {
            margin: 40px 0;
            padding: 30px;
            background: rgba(255, 255, 255, 0.05);
            border-radius: 15px;
            border: 1px solid rgba(255, 255, 255, 0.1);
        }
        
        .blog-section h2 {
            font-size: 2em;
            margin-bottom: 25px;
            background: linear-gradient(45deg, #4ecdc4, #45b7d1);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text;
            text-align: center;
        }
        
        .blog-content p {
            font-size: 1.1em;
            line-height: 1.8;
            margin-bottom: 20px;
            opacity: 0.95;
            text-align: left;
        }
        
        .blog-content p:last-child {
            margin-bottom: 0;
            font-style: italic;
            opacity: 0.9;
        }
        
        .footer {
            margin-top: 40px;
            font-size: 0.9em;
            opacity: 0.7;
        }
        
        @media (max-width: 600px) {
            .container {
                margin: 20px;
                padding: 30px 20px;
            }
            
            h1 {
                font-size: 2.2em;
            }
            
            .subtitle {
                font-size: 1.2em;
            }
            
            .description {
                font-size: 1em;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="profile-pic">P</div>
        <h1>Hi, I'm Purnachandra!</h1>
        <div class="subtitle">Recent Master's Graduate</div>
        <div class="description">
            I'm passionate about cloud technologies and DevOps tools that make life easier for developers and teams. 
            I love exploring how modern infrastructure and automation can streamline workflows and improve productivity.
        </div>
        
        <div class="tech-stack">
            <span class="tech-tag">☁️ Cloud Computing</span>
            <span class="tech-tag">🐳 Docker</span>
            <span class="tech-tag">⚙️ DevOps</span>
            <span class="tech-tag">🚀 Automation</span>
            <span class="tech-tag">🔧 Infrastructure</span>
        </div>
        
        <div class="description">
            Let's connect if we're on the same page! I'm always excited to discuss cloud technologies, 
            share knowledge, and collaborate on interesting projects.
        </div>
        
        <div class="blog-section">
            <h2>Beyond the Code</h2>
            <div class="blog-content">
                <p>When I'm not diving deep into cloud technologies, you'll find me exploring the world and embracing life's adventures. I'm an extrovert who thrives on meaningful connections and enjoys spending quality time with people from all walks of life.</p>
                
                <p>Traveling is my passion – there's something magical about discovering new cultures, tasting local cuisines, and creating memories in places I've never been. Each journey teaches me something new and broadens my perspective on life.</p>
                
                <p>Fitness is a cornerstone of my daily routine. Whether I'm hitting the gym for strength training or going for a peaceful run in the morning, these activities bring me balance and inner peace. There's nothing quite like the clarity that comes from a good workout or a refreshing jog.</p>
                
                <p>I believe in living life to the fullest – connecting with people, staying active, and finding joy in both the big adventures and the simple moments. Life is about balance, growth, and making meaningful connections along the way.</p>
            </div>
        </div>
        
        <a href="mailto:purnachandra@example.com" class="connect-btn">Let's Connect!</a>
        
        <div class="footer">
            <p>🚀 Hosted with Docker & Go | Made with ❤️</p>
        </div>
    </div>
</body>
</html>
		`)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
